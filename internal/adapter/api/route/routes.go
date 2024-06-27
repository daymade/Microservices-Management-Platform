package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/middleware"
)

func SetupRoutes(r *gin.Engine, sh *handler.ServiceHandler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth()) // 使用认证中间件

	v1.GET("/services", sh.ListServices)
	v1.GET("/services/:id", sh.GetService)
}
