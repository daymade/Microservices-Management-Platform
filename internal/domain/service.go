package domain

import (
	"catalog-service-management-api/internal/domain/models"
)

type Service interface {
	ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error)
	GetService(id string) (models.Service, error)
}
