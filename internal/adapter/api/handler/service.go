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
