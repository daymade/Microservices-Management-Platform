package storage

import (
	"catalog-service-management-api/internal/domain/models"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemoryStorage_ListServices(t *testing.T) {
	// 创建测试数据
	now := time.Now()
	services := []models.Service{
		{ID: "1", Name: "Service A", Description: "Description A", CreatedAt: now.Add(-3 * time.Hour)},
		{ID: "2", Name: "Service B", Description: "Description B", CreatedAt: now.Add(-2 * time.Hour)},
		{ID: "3", Name: "Service C", Description: "Description C", CreatedAt: now.Add(-1 * time.Hour)},
		{ID: "4", Name: "Another Service", Description: "Another Description", CreatedAt: now},
	}

	m := make(map[string]models.Service)
	for _, s := range services {
		m[s.ID] = s
	}

	storage := &MemoryStorage{
		services: m,
		mutex:    sync.RWMutex{},
	}

	tests := []struct {
		name     string
		query    string
		sortBy   string
		sortDir  string
		page     int
		pageSize int
		expected []models.Service
		total    int
	}{
		{
			name:     "List all services, default sort",
			query:    "",
			sortBy:   "",
			sortDir:  "",
			page:     1,
			pageSize: 10,
			expected: []models.Service{services[3], services[2], services[1], services[0]},
			total:    4,
		},
		{
			name:     "Filter services by name",
			query:    "service a",
			sortBy:   "",
			sortDir:  "",
			page:     1,
			pageSize: 10,
			expected: []models.Service{services[0]},
			total:    1,
		},
		{
			name:     "Sort by name ascending",
			query:    "",
			sortBy:   "name",
			sortDir:  "asc",
			page:     1,
			pageSize: 10,
			expected: []models.Service{services[3], services[0], services[1], services[2]},
			total:    4,
		},
		{
			name:     "Sort by name descending",
			query:    "",
			sortBy:   "name",
			sortDir:  "desc",
			page:     1,
			pageSize: 10,
			expected: []models.Service{services[2], services[1], services[0], services[3]},
			total:    4,
		},
		{
			name:     "Sort by created_at ascending",
			query:    "",
			sortBy:   "created_at",
			sortDir:  "asc",
			page:     1,
			pageSize: 10,
			expected: []models.Service{services[0], services[1], services[2], services[3]},
			total:    4,
		},
		{
			name:     "Sort by created_at descending",
			query:    "",
			sortBy:   "created_at",
			sortDir:  "desc",
			page:     1,
			pageSize: 10,
			expected: []models.Service{services[3], services[2], services[1], services[0]},
			total:    4,
		},
		{
			name:     "Pagination - first page",
			query:    "",
			sortBy:   "",
			sortDir:  "",
			page:     1,
			pageSize: 2,
			expected: []models.Service{services[3], services[2]},
			total:    4,
		},
		{
			name:     "Pagination - second page",
			query:    "",
			sortBy:   "",
			sortDir:  "",
			page:     2,
			pageSize: 2,
			expected: []models.Service{services[1], services[0]},
			total:    4,
		},
		{
			name:     "Pagination - out of range",
			query:    "",
			sortBy:   "",
			sortDir:  "",
			page:     3,
			pageSize: 2,
			expected: []models.Service{},
			total:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, total, err := storage.ListServices(tt.query, tt.sortBy, tt.sortDir, tt.page, tt.pageSize)
			assert.NoError(t, err)
			assert.Equal(t, tt.total, total)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMemoryStorage_GetService(t *testing.T) {
	// 创建测试数据
	now := time.Now()
	services := []models.Service{
		{ID: "1", Name: "Service A", Description: "Description A", CreatedAt: now.Add(-3 * time.Hour)},
		{ID: "2", Name: "Service B", Description: "Description B", CreatedAt: now.Add(-2 * time.Hour)},
		{ID: "3", Name: "Service C", Description: "Description C", CreatedAt: now.Add(-1 * time.Hour)},
		{ID: "4", Name: "Another Service", Description: "Another Description", CreatedAt: now},
	}

	m := make(map[string]models.Service)
	for _, s := range services {
		m[s.ID] = s
	}

	storage := &MemoryStorage{
		services: m,
		mutex:    sync.RWMutex{},
	}

	tests := []struct {
		name        string
		id          string
		expected    models.Service
		expectedErr bool
	}{
		{
			name:        "Get existing service",
			id:          "1",
			expected:    services[0],
			expectedErr: false,
		},
		{
			name:        "Get another existing service",
			id:          "4",
			expected:    services[3],
			expectedErr: false,
		},
		{
			name:        "Get non-existent service",
			id:          "5",
			expected:    models.Service{},
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := storage.GetService(tt.id)

			if tt.expectedErr {
				assert.Error(t, err)
				assert.Equal(t, "service not found", err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
