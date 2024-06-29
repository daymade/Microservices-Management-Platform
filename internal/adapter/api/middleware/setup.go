package middleware

import (
	"github.com/gin-gonic/gin"
)

// SetupGlobalMiddleware 配置所有全局中间件
func SetupGlobalMiddleware(r *gin.Engine, appName string) {
	// 使用 tracing 中间件
	r.Use(NewTracingMiddleware(appName))

	// HTTP Metrics 中间件配置
	r.Use(HttpMetrics(appName))

	// CORS 中间件配置
	r.Use(Cors())
}
