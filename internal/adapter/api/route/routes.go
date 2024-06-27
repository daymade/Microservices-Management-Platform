package route

import (
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/middleware"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, sh *handler.ServiceHandler, uh *handler.UserHandler) {
	// Expose the registered metrics at `/metrics` path.
	r.GET("/metrics", func(c *gin.Context) {
		metrics.WritePrometheus(c.Writer, true)
	})

	// swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Metrics 中间件配置
	r.Use(middleware.Metrics())

	// CORS 中间件配置
	r.Use(middleware.Cors())

	// API v1 路由组
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth()) // 使用认证中间件

	// 服务相关路由
	v1.GET("/services", sh.ListServices)
	v1.GET("/services/:id", sh.GetService)

	// 用户相关路由
	v1.GET("/user", uh.GetCurrentUser)
}
