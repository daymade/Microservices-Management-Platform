package route

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/middleware"
)

func SetupRoutes(engine *gin.Engine, sh *handler.ServiceHandler) {
	v1 := engine.Group("/v1")
	v1.Use(middleware.Auth()) // 使用认证中间件

	v1.GET("/services", sh.ListServices)
	v1.GET("/services/:id", sh.GetService)
}
