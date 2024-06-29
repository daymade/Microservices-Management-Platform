package route

import (
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, sh *handler.ServiceHandler, uh *handler.UserHandler) {
	// swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 路由组
	v1 := r.Group("/api/v1")

	// 只为 v1 路由组设置认证中间件
	v1.Use(middleware.Auth())

	// 服务相关路由
	setupServiceRoutes(v1, sh)

	// 用户相关路由
	setupUserRoutes(v1, uh)
}

func setupServiceRoutes(rg *gin.RouterGroup, sh *handler.ServiceHandler) {
	rg.GET("/services", sh.ListServices)
	rg.GET("/services/:id", sh.GetService)
}

func setupUserRoutes(rg *gin.RouterGroup, uh *handler.UserHandler) {
	rg.GET("/user", uh.GetCurrentUser)
}
