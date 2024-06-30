package entity

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"index"`
}

type Service struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"not null;index"`
	Description string
	OwnerID     uint64
	Owner       User      `gorm:"foreignKey:OwnerID"`
	Versions    []Version `gorm:"foreignKey:ServiceID"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time
}

type Version struct {
	ID          uint64  `gorm:"primaryKey"`
	ServiceID   uint64  `gorm:"index"`
	Service     Service `gorm:"foreignKey:ServiceID"`
	Number      string  `gorm:"not null"`
	Description string
	CreatedAt   time.Time `gorm:"index"`
}
