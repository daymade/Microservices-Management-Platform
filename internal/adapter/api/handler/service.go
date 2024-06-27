package handler

import (
	"errors"
	"catalog-service-management-api/internal/adapter/api/viewmodel"
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
// @Summary      列出服务
// @Description  获取服务列表，支持分页、排序和搜索
// @Tags         services
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        query           query     string  false  "搜索关键词"
// @Param        sort_by         query     string  false  "排序字段"
// @Param        sort_direction  query     string  false  "排序方向 (asc 或 desc)"
// @Param        page            query     int     false  "页码"
// @Param        page_size       query     int     false  "每页条数"
// @Success      200  {object}  viewmodel.PaginatedResponse{data=[]viewmodel.ServiceListViewModel}
// @Failure      400  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /services [get]
func (h *ServiceHandler) ListServices(c *gin.Context) {
	query := c.DefaultQuery("query", "")
	sortBy := c.DefaultQuery("sort_by", "name")
	sortDir := c.DefaultQuery("sort_direction", "asc")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// Validate input
	if err := validateListServicesInput(query, sortBy, sortDir, page, pageSize); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services, total, err := h.manager.ListServices(query, sortBy, sortDir, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to view models
	svm := make([]viewmodel.ServiceListViewModel, len(services))
	for i, s := range services {
		svm[i] = viewmodel.NewServiceListViewModel(s)
	}

	response := viewmodel.NewPaginatedResponse(svm, total, page, pageSize)
	c.JSON(http.StatusOK, response)
}

// GetService godoc
// @Summary      获取单个服务
// @Description  通过 ID 获取单个服务的详细信息
// @Tags         services
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        id   path      string  true  "服务 ID"
// @Success      200  {object}  viewmodel.ServiceDetailViewModel
// @Failure      400  {object}  map[string]any
// @Failure      404  {object}  map[string]any
// @Router       /services/{id} [get]
func (h *ServiceHandler) GetService(c *gin.Context) {
	id := c.Param("id")

	// Validate input
	if err := validateServiceID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := h.manager.GetService(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	svm := viewmodel.NewServiceDetailViewModel(s)

	c.JSON(http.StatusOK, svm)
}

func validateListServicesInput(query, sortBy, sortDir string, page, pageSize int) error {
	if page < 1 {
		return errors.New("page must be greater than 0")
	}
	if pageSize < 1 || pageSize > 100 {
		return errors.New("page_size must be between 1 and 100")
	}
	if sortDir != "asc" && sortDir != "desc" {
		return errors.New("sort_direction must be 'asc' or 'desc'")
	}
	// Add more validations as needed
	return nil
}

func validateServiceID(id string) error {
	// Add your validation logic here
	if id == "" {
		return errors.New("service ID cannot be empty")
	}
	return nil
}
