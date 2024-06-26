package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Avatar    string    `json:"avatar"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
