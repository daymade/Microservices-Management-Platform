package viewmodel

import (
	"catalog-service-management-api/internal/domain/models"
	"time"
)

func NewServiceListViewModel(s models.Service) ServiceListViewModel {
	return ServiceListViewModel{
		ID:           s.ID,
		Name:         s.Name,
		Description:  s.Description,
		OwnerID:      s.OwnerID,
		VersionCount: len(s.Versions),
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func NewServiceDetailViewModel(s models.Service) ServiceDetailViewModel {
	versions := make([]VersionViewModel, len(s.Versions))
	for i, v := range s.Versions {
		versions[i] = VersionViewModel{
			Number:      v.Number,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
		}
	}

	return ServiceDetailViewModel{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		OwnerID:     s.OwnerID,
		Versions:    versions,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

// ServiceListViewModel 用于服务列表展示的视图模型
// @Description 服务列表项模型，包含服务的基本信息
type ServiceListViewModel struct {
	ID           string    `json:"id" example:"srv-123"`
	Name         string    `json:"name" example:"My Service"`
	Description  string    `json:"description" example:"This is a sample service"`
	OwnerID      string    `json:"owner_id" example:"usr-456"`
	VersionCount int       `json:"version_count" example:"3"`
	CreatedAt    time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-01-02T00:00:00Z"`
}

// ServiceDetailViewModel 用于服务详情展示的视图模型
// @Description 服务详情模型，包含服务的详细信息和版本列表
type ServiceDetailViewModel struct {
	ID          string             `json:"id" example:"srv-123"`
	Name        string             `json:"name" example:"My Service"`
	Description string             `json:"description" example:"This is a sample service"`
	OwnerID     string             `json:"owner_id" example:"usr-456"`
	Versions    []VersionViewModel `json:"versions"`
	CreatedAt   time.Time          `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time          `json:"updated_at" example:"2023-01-02T00:00:00Z"`
}

// VersionViewModel 用于版本信息展示的视图模型
// @Description 版本信息模型，包含版本的基本信息
type VersionViewModel struct {
	Number      string    `json:"number" example:"v1.0.0"`
	Description string    `json:"description" example:"Initial release"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
}
