package storage

import (
	"fmt"
	"catalog-service-management-api/internal/domain/models"
	"math/rand"
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
		mutex:    sync.RWMutex{},
	}

	if len(storage.services) == 0 {
		// 随机填充一些测试数据
		count := 25
		rand.Seed(time.Now().UnixNano())

		baseTime := time.Now().Add(-365 * 24 * time.Hour) // Start from one year ago
		for i := 1; i <= count; i++ {
			id := fmt.Sprintf("%d", i)
			storage.services[id] = generateRandomService(id, baseTime)
			baseTime = baseTime.Add(time.Duration(rand.Intn(24)) * time.Hour) // Add up to 24 hours
		}
	}

	return storage
}

func (m *MemoryStorage) ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	// 过滤
	var filteredServices []models.Service
	for _, service := range m.services {
		if query == "" ||
			strings.Contains(strings.ToLower(service.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(service.Description), strings.ToLower(query)) {
			filteredServices = append(filteredServices, service)
		}
	}

	// 排序（每个 Service 的 Version 排序在 domain/app 层处理）
	switch sortBy {
	case "name":
		sort.Slice(filteredServices, func(i, j int) bool {
			if sortDir == "asc" {
				return filteredServices[i].Name < filteredServices[j].Name
			}
			return filteredServices[i].Name > filteredServices[j].Name
		})
	case "created_at":
		sort.Slice(filteredServices, func(i, j int) bool {
			if sortDir == "asc" {
				return filteredServices[i].CreatedAt.Before(filteredServices[j].CreatedAt)
			}
			return filteredServices[i].CreatedAt.After(filteredServices[j].CreatedAt)
		})
	default:
		// 默认按创建时间降序排序
		sort.Slice(filteredServices, func(i, j int) bool {
			return filteredServices[i].CreatedAt.After(filteredServices[j].CreatedAt)
		})
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

// 随便找点单词自动生成测试数据
var (
	adjectives = []string{"Advanced", "Intelligent", "Secure", "Scalable", "Efficient", "Innovative", "Robust", "Flexible", "Integrated", "Automated"}
	nouns      = []string{"Analytics", "Database", "Networking", "Messaging", "Authentication", "Monitoring", "Deployment", "Testing", "Backup", "Encryption"}
	verbs      = []string{"optimizes", "enhances", "secures", "streamlines", "automates", "simplifies", "accelerates", "manages", "analyzes", "transforms"}
)

func generateRandomService(id string, baseTime time.Time) models.Service {
	name := fmt.Sprintf("%s %s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))])
	description := fmt.Sprintf("This service %s your %s processes", verbs[rand.Intn(len(verbs))], strings.ToLower(nouns[rand.Intn(len(nouns))]))

	createdAt := baseTime
	updatedAt := createdAt.Add(time.Duration(rand.Intn(int(time.Since(createdAt).Hours()))) * time.Hour)

	versionCount := rand.Intn(5) + 1 // 1 to 5 versions
	versions := make([]models.Version, versionCount)
	for i := 0; i < versionCount; i++ {
		versions[i] = models.Version{
			Number:      fmt.Sprintf("v1.%d", i),
			Description: fmt.Sprintf("Version %d of the service", i+1),
			CreatedAt:   createdAt.Add(time.Duration(i*30*24) * time.Hour), // Each version is about a month apart
		}
	}

	return models.Service{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     fmt.Sprintf("user%d", rand.Intn(10)+1),
		Versions:    versions,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
