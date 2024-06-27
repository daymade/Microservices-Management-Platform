package handler

import (
	"catalog-service-management-api/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	manager *service.Manager
}

func NewServiceHandler(m *service.Manager) *ServiceHandler {
	return &ServiceHandler{manager: m}
}

// ListServices godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "search by q"
// @Param        sort_by    query     string  false  "sort by field"
// @Param        sort_direction    query     string  false  "sort direction"
// @Param        page    query     int  false  "page number"
// @Param        page_size    query     int  false  "page size"
// @Success      200  {object}  viewmodel.Service
// @Router       /services [get]
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
