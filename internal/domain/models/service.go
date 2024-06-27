package models

import "time"

// Service 表示一个服务实体
// 它包含了服务的基本信息、版本历史以及创建和更新时间
type Service struct {
	// ID 是服务的唯一标识符
	ID string `json:"id"`

	// Name 是服务的名称
	Name string `json:"name"`

	// Description 是服务的详细描述
	Description string `json:"description"`

	// OwnerID 是服务所有者的唯一标识符，关联到 User 实体
	OwnerID string `json:"owner_id"`

	// Versions 是该服务的所有版本列表
	Versions []Version `json:"versions"`

	// CreatedAt 表示服务创建的时间
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 表示服务最后更新的时间
	UpdatedAt time.Time `json:"updated_at"`
}

// Version 表示服务的一个特定版本
// 它包含版本号、描述和创建时间
type Version struct {
	// Number 是版本号，通常遵循语义化版本规范（如 v1.0.0）
	Number string `json:"number"`

	// Description 是该版本的详细描述，通常包括更新内容或变更说明
	Description string `json:"description"`

	// CreatedAt 表示该版本创建的时间
	CreatedAt time.Time `json:"created_at"`
}
