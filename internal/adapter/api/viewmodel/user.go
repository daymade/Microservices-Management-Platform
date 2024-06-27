package viewmodel

import "time"

// User 表示用户信息的视图模型
// @Description 用户信息模型，包含用户的基本信息
type User struct {
	ID        string    `json:"id" example:"1"`
	Avatar    string    `json:"avatar" example:"https://example.com/avatar.jpg"`
	Username  string    `json:"username" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
}
