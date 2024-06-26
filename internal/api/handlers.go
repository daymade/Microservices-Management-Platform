package api

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/storage"
	"net/http"
	"strconv"
)

type Handler struct {
	storage *storage.MemoryStorage
}

func NewHandler(storage *storage.MemoryStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) ListServices(c *gin.Context) {
	query := c.DefaultQuery("query", "")
	sortBy := c.DefaultQuery("sort_by", "name")
	sortDir := c.DefaultQuery("sort_direction", "asc")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	services, total, err := h.storage.ListServices(query, sortBy, sortDir, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"services":    services,
		"total_count": total,
		"page":        page,
		"page_size":   pageSize,
	})
}

func (h *Handler) GetService(c *gin.Context) {
	id := c.Param("id")

	service, err := h.storage.GetService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}
