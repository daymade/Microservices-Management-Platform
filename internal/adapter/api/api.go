package api

import (
	"context"
	"errors"
	"fmt"
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/middleware"
	"catalog-service-management-api/internal/adapter/api/route"
	"catalog-service-management-api/internal/app/service"
	"catalog-service-management-api/internal/infrastructure/otel/meter"
	"catalog-service-management-api/internal/infrastructure/otel/tracing"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	appName = "catalog-service-management"
	env     = "testing"
)

type API struct {
}

func NewHTTPServer() *API {
	return &API{}
}

func (a *API) Run(addr string) error {
	r := gin.Default()

	// 设置全局中间件
	middleware.SetupGlobalMiddleware(r, appName)

	sh := handler.NewServiceHandler(service.NewManager())
	uh := handler.NewUserHandler()

	route.SetupRoutes(r, sh, uh)

	return r.Run(addr)
}

func (a *API) RunOTel(ctx context.Context, metricsAddr string) error {
	// 初始化 tracer
	tracerConfig := tracing.Config{
		AppName:        appName,
		Env:            env,
		JaegerEndpoint: "http://jaeger:14268/api/traces", // 建议从配置文件或环境变量中读取
		CheckInterval:  10 * time.Second,
	}
	tracer := tracing.NewTracer(tracerConfig)
	tracer.Init(ctx)

	defer func() {
		if err := tracer.Shutdown(context.Background()); err != nil {
			log.Printf("关闭tracer provider时出错: %v", err)
		}
	}()

	// 初始化 meter
	meterConfig := meter.Config{
		AppName:           appName,
		Env:               env,
		CollectorEndpoint: "otel-collector:4317",
	}
	mp, err := meter.InitMeter(ctx, meterConfig)
	if err != nil {
		return fmt.Errorf("初始化 meter 失败: %w", err)
	}

	defer func() {
		if err := mp.Shutdown(context.Background()); err != nil {
			log.Printf("关闭 meter provider 时出错: %v", err)
		}
	}()

	// 使用 gin.New() 而不是 gin.Default() 以避免默认的日志中间件
	metricsRouter := gin.New()

	metricsRouter.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Printf("metrics 监听地址: %s", metricsAddr)

	server := &http.Server{
		Addr:    metricsAddr,
		Handler: metricsRouter,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("metrics 服务器运行错误: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}
