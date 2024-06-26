package viewmodel

import "time"

type Service struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerID      string    `json:"owner_id"`
	VersionCount int       `json:"version_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ServiceDetail struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	Versions    []Version `json:"versions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Version struct {
	Number      string `json:"number"`
	Description string `json:"description"`
}
