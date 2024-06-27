package route

import (
	"catalog-service-management-api/internal/adapter/api/handler"
	"catalog-service-management-api/internal/adapter/api/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, sh *handler.ServiceHandler, uh *handler.UserHandler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORS 中间件配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth()) // 使用认证中间件

	// 服务相关路由
	v1.GET("/services", sh.ListServices)
	v1.GET("/services/:id", sh.GetService)

	// 用户相关路由
	v1.GET("/user", uh.GetCurrentUser)
}
