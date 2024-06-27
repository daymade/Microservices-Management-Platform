package models

import "time"

// User 表示系统中的用户实体
// 它包含了用户的基本信息，如ID、头像、用户名、邮箱和创建时间
type User struct {
	// ID 是用户的唯一标识符，数据库中是 BIGINT 类型，为了兼容前端 js，这里使用字符串类型
	ID string `json:"id"`

	// Avatar 是用户头像的URL
	// 如果用户没有设置头像，这可能是一个空字符串或默认头像的URL, 前端会处理 avatar 为空的情况
	Avatar string `json:"avatar"`

	// Username 是用户的显示名称
	Username string `json:"username"`

	// Email 是用户的电子邮箱地址
	Email string `json:"email"`

	// CreatedAt 表示用户账户创建的时间
	CreatedAt time.Time `json:"created_at"`
}
