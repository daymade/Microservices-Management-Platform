package storage

import (
	"fmt"
	"catalog-service-management-api/internal/domain/models"
	"sort"
	"strings"
	"sync"
	"time"
)

type MemoryStorage struct {
	services map[string]models.Service
	mutex    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	storage := &MemoryStorage{
		services: make(map[string]models.Service),
	}

	storage.services["1"] = models.Service{
		ID:          "1",
		Name:        "Service One",
		Description: "This is the first service",
		OwnerID:     "user1",
		Versions: []models.Version{
			{Number: "v1.0", Description: "Initial version", CreatedAt: time.Now()},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	storage.services["2"] = models.Service{
		ID:          "2",
		Name:        "Service Two",
		Description: "This is the second service",
		OwnerID:     "user2",
		Versions: []models.Version{
			{Number: "v1.0", Description: "Initial version", CreatedAt: time.Now()},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return storage
}

func (m *MemoryStorage) ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var filteredServices []models.Service
	for _, service := range m.services {
		if query == "" || strings.Contains(strings.ToLower(service.Name), strings.ToLower(query)) || strings.Contains(strings.ToLower(service.Description), strings.ToLower(query)) {
			filteredServices = append(filteredServices, service)
		}
	}

	// 排序
	if sortBy == "name" {
		if sortDir == "asc" {
			sort.Slice(filteredServices, func(i, j int) bool {
				return filteredServices[i].Name < filteredServices[j].Name
			})
		} else {
			sort.Slice(filteredServices, func(i, j int) bool {
				return filteredServices[i].Name > filteredServices[j].Name
			})
		}
	}

	// 分页
	total := len(filteredServices)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		return []models.Service{}, total, nil
	}
	if end > total {
		end = total
	}

	return filteredServices[start:end], total, nil
}

func (m *MemoryStorage) GetService(id string) (models.Service, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	service, exists := m.services[id]
	if !exists {
		return models.Service{}, fmt.Errorf("service not found")
	}

	return service, nil
}
