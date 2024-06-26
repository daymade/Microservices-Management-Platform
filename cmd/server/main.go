package main

import (
	"fmt"
	"catalog-service-management-api/internal/adapter/api"
	"os"
	"os/signal"
	"syscall"
)

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
