package handler

import (
	"github.com/gin-gonic/gin"
	"catalog-service-management-api/internal/app/service"
	"net/http"
	"strconv"
)

type ServiceHandler struct {
	manager *service.Manager
}

func NewServiceHandler(m *service.Manager) *ServiceHandler {
	return &ServiceHandler{manager: m}
}

func (h *ServiceHandler) ListServices(c *gin.Context) {
	query := c.DefaultQuery("query", "")
	sortBy := c.DefaultQuery("sort_by", "name")
	sortDir := c.DefaultQuery("sort_direction", "asc")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	services, total, err := h.manager.ListServices(query, sortBy, sortDir, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	service_view_models := services

	c.JSON(http.StatusOK, gin.H{
		"services":    service_view_models,
		"total_count": total,
		"page":        page,
		"page_size":   pageSize,
	})
}

func (h *ServiceHandler) GetService(c *gin.Context) {
	id := c.Param("id")

	s, err := h.manager.GetService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	service_view_model := s

	c.JSON(http.StatusOK, service_view_model)
}
