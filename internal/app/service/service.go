package service

import (
	"fmt"
	"catalog-service-management-api/internal/app/util/version"
	"catalog-service-management-api/internal/domain"
	"catalog-service-management-api/internal/domain/models"
	"catalog-service-management-api/internal/infrastructure/storage"
	"log"
	"os"
)

// Manager 定义了服务管理的类,
// app 层依赖 domain 层的接口，domain 的接口由 infra 层实现，app 负责注入 infra 到 domain，依赖关系为：app -> domain <- infra
type Manager struct {
	service domain.Service
}

// NewManager 创建一个新的 Manager 实例
func NewManager() *Manager {
	var s domain.Service
	var err error

	// 是否使用外部数据库
	useDB := os.Getenv("USE_DB") == "true"
	// 选择存储实现
	if useDB {
		s, err = storage.NewPostgresStorage()
		if err != nil {
			log.Fatalln("Failed to connect to database:", err)
		}
	} else {
		s = storage.NewMemoryStorage()
	}

	return &Manager{
		service: s,
	}
}

// ListServices 列出服务
// Demo 的功能比较简单，所以 app 层比较薄，实际项目中可能会有更多的业务逻辑，比如聚合多个 domain 层的 service
func (sm *Manager) ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error) {
	services, total, err := sm.service.ListServices(query, sortBy, sortDir, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 对每个服务的版本进行排序
	for i := range services {
		services[i].Versions = version.Sort(services[i].Versions)
	}

	return services, total, nil
}

// GetService 获取单个服务
func (sm *Manager) GetService(id string) (models.Service, error) {
	service, err := sm.service.GetService(id)
	if err != nil {
		return models.Service{}, err
	}

	// 对服务的版本进行排序
	service.Versions = version.Sort(service.Versions)

	return service, nil
}

// GetVersions 获取服务版本
func (sm *Manager) GetVersions(id string) ([]models.Version, error) {
	service, err := sm.service.GetService(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	// 排序
	versions := version.Sort(service.Versions)

	return versions, nil
}
