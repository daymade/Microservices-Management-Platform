package main

import (
	"fmt"
	_ "catalog-service-management-api/api"
	"catalog-service-management-api/internal/adapter/api"
	"os"
	"os/signal"
	"syscall"
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

// @securityDefinitions.basic  BasicAuth
func main() {
	http := api.NewHTTPServer()

	// 启动 http api server
	go func() {
		if err := http.Run(":8080"); err != nil {
			fmt.Println("Failed to run server:", err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// 等待信号
	<-quit
}
