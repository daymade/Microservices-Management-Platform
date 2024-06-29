package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer     trace.Tracer
	tracerLock sync.RWMutex
)

const (
	GinErrorKey              = "gin.errors"
	ResponseHeaderTraceIdKey = "X-Trace-ID" // 定义响应头的名称
	GinContextTraceIdKey     = "X-Jaeger-Trace-ID"
)

type TracingMiddleware struct {
	appName string
}

func NewTracingMiddleware(appName string) gin.HandlerFunc {
	tm := &TracingMiddleware{
		appName: appName,
	}

	go tm.periodicCheck()

	return tm.handler()
}

func (tm *TracingMiddleware) periodicCheck() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		tm.updateTracer()
	}
}

func (tm *TracingMiddleware) updateTracer() {
	tp := otel.GetTracerProvider()
	if tp != nil {
		newTracer := tp.Tracer(tm.appName)
		tracerLock.Lock()
		tracer = newTracer
		tracerLock.Unlock()
	}
}

func (tm *TracingMiddleware) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracerLock.RLock()
		currentTracer := tracer
		tracerLock.RUnlock()

		if currentTracer == nil {
			c.Next()
			return
		}

		propagator := otel.GetTextMapPropagator()
		ctx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

		spanName := c.FullPath()
		if spanName == "" {
			spanName = c.Request.URL.Path
		}

		ctx, span := currentTracer.Start(
			ctx,
			spanName,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				semconv.HTTPMethodKey.String(c.Request.Method),
				semconv.HTTPRouteKey.String(c.FullPath()),
				semconv.HTTPSchemeKey.String(c.Request.URL.Scheme),
				semconv.HTTPTargetKey.String(c.Request.URL.Path),
				semconv.NetHostNameKey.String(c.Request.Host),
				semconv.HTTPFlavorKey.String(c.Request.Proto),
				semconv.HTTPRequestContentLengthKey.Int64(c.Request.ContentLength),
			),
		)
		defer span.End()

		// 获取 TraceID 并将其添加到 Context 中
		traceID := span.SpanContext().TraceID().String()
		c.Set(GinContextTraceIdKey, traceID)

		// 在响应头中设置 TraceID
		c.Header(ResponseHeaderTraceIdKey, traceID)

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		status := c.Writer.Status()
		attrs := []attribute.KeyValue{
			semconv.HTTPStatusCodeKey.Int(status),
		}

		if len(c.Errors) > 0 {
			span.SetAttributes(attribute.String(GinErrorKey, c.Errors.String()))
		}

		span.SetAttributes(attrs...)
	}
}
