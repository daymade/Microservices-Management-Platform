package api

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/storage"
	"net/http"
)

type Handler struct {
	storage *storage.MemoryStorage
}

func NewHandler(storage *storage.MemoryStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) ListServices(c *gin.Context) {
	// 实现列表服务的处理逻辑
	c.JSON(http.StatusOK, gin.H{"message": "ListServices endpoint"})
}

func (h *Handler) GetService(c *gin.Context) {
	// 实现获取单个服务的处理逻辑
	c.JSON(http.StatusOK, gin.H{"message": "GetService endpoint"})
}
