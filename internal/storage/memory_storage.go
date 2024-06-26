package storage

import (
	"catalog-service-management-api/internal/models"
	"sync"
)

type MemoryStorage struct {
	services map[string]models.Service
	mutex    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		services: make(map[string]models.Service),
	}
}

func (m *MemoryStorage) ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error) {
	// 实现列表、搜索、排序和分页逻辑
	return nil, 0, nil
}

func (m *MemoryStorage) GetService(id string) (models.Service, error) {
	// 实现获取单个服务的逻辑
	return models.Service{}, nil
}
