package main

import (
	"context"
	"fmt"
	// _ imports the swagger docs
	_ "catalog-service-management-api/api"
	"catalog-service-management-api/internal/adapter/api"
	"os"
	"os/signal"
	"syscall"
)

const (
	srvAddr     = ":8080"
	metricsAddr = ":9090"
)

// @title           Catalog Service Management
// @version         1.0
// @description     This is a platform to manage services.

// @contact.name   daymade
// @contact.email  daymadev89@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http := api.NewHTTPServer()

	// 启动 http api server
	go func() {
		if err := http.Run(srvAddr); err != nil {
			fmt.Println("Failed to run server:", err)
		}
	}()

	// 启动 open telemetry 的 metrics 和 tracing
	go func() {
		if err := http.RunOTel(ctx, metricsAddr); err != nil {
			fmt.Println("Failed to run metrics server:", err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 等待信号
	<-quit
}
