package entity

import (
	"time"
)

type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
}

type Service struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	OwnerID     uint64
	Owner       User `gorm:"foreignKey:OwnerID"`
	Versions    []Version
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Version struct {
	ID          uint64 `gorm:"primaryKey"`
	ServiceID   uint64
	Service     Service `gorm:"foreignKey:ServiceID"`
	Number      string  `gorm:"not null"`
	Description string
	CreatedAt   time.Time
}
