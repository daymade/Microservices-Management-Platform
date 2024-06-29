package tracing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

const envKey = "environment"

type Config struct {
	AppName        string
	Env            string
	JaegerEndpoint string
	CheckInterval  time.Duration
}

type Tracer struct {
	config Config
	tp     *trace.TracerProvider
	tpMu   sync.RWMutex
}

func NewTracer(config Config) *Tracer {
	t := &Tracer{
		config: config,
		tpMu:   sync.RWMutex{},
	}
	return t
}

func (t *Tracer) Init(ctx context.Context) {
	// TODO 现在应用启动后没有马上启动 tracer，可能会导致一些请求没有被追踪
	go t.checkAndInitTracer(ctx)
}

func (t *Tracer) checkAndInitTracer(ctx context.Context) {
	ticker := time.NewTicker(t.config.CheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			t.tpMu.Lock()
			if t.isJaegerAvailable() {
				if t.tp == nil {
					if err := t.initializeTracer(); err != nil {
						log.Printf("初始化 tracer 失败: %v", err)
					} else {
						log.Printf("使用 Jaeger 初始化 tracing 成功")
					}
				}
			} else {
				if t.tp != nil {
					log.Printf("Jaeger 不可用，关闭当前 TracerProvider")
					if err := t.tp.Shutdown(ctx); err != nil {
						log.Printf("关闭 TracerProvider 失败: %v", err)
					}
					t.tp = nil
				}
			}
			t.tpMu.Unlock()
		}
	}
}

func (t *Tracer) isJaegerAvailable() bool {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(t.config.JaegerEndpoint)
	if err != nil {
		log.Printf("Jaeger 不可用: %v", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusMethodNotAllowed {
		log.Printf("Jaeger 暂时不可用: %d，跳过 tracing 直到 Jaeger 可用", resp.StatusCode)
		return false
	}
	return true
}

func (t *Tracer) initializeTracer() error {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(t.config.JaegerEndpoint)))
	if err != nil {
		return fmt.Errorf("创建 Jaeger 导出器失败: %w", err)
	}

	t.tp = trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(t.config.AppName),
			attribute.String(envKey, t.config.Env),
		)),
	)

	// 这里设置全局变量不是很优雅，检查 tracer 可用和恢复的代码分别在 api/middleware 和 infra 里面
	// 理想情况下 api/middleware 里只需要调用 tracing 的方法，不应该负责检查 tracer 的可用不可用
	otel.SetTracerProvider(t.tp)
	return nil
}

func (t *Tracer) Shutdown(ctx context.Context) error {
	t.tpMu.RLock()
	defer t.tpMu.RUnlock()
	if t.tp != nil {
		return t.tp.Shutdown(ctx)
	}
	return nil
}
