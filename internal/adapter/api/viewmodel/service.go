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

// ServiceListViewModel 用于列表展示
type ServiceListViewModel struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerID      string    `json:"owner_id"`
	VersionCount int       `json:"version_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ServiceDetailViewModel 用于详情展示
type ServiceDetailViewModel struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	OwnerID     string             `json:"owner_id"`
	Versions    []VersionViewModel `json:"versions"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type VersionViewModel struct {
	Number      string    `json:"number"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
