package models

import "time"

type Service struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	Versions    []Version `json:"versions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Version struct {
	Number      string    `json:"number"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
