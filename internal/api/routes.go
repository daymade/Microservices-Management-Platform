package api

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/middleware"
)

func SetupRoutes(r *gin.Engine, h *Handler) {
	v1 := r.Group("/v1")
	v1.Use(middleware.Auth()) // 使用认证中间件

	v1.GET("/services", h.ListServices)
	v1.GET("/services/:id", h.GetService)
}
