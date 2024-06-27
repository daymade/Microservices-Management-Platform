package handler

import (
	"catalog-service-management-api/internal/adapter/api/viewmodel"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UserHandler 处理用户相关的请求
type UserHandler struct {
	// 可以添加用户服务等依赖
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetCurrentUser 获取当前登录用户的信息
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	// 这里应该从认证中间件中获取用户ID
	// 为了演示，我们直接返回一个模拟的用户
	user := viewmodel.User{
		ID:        "1",
		Avatar:    "",
		Username:  "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, user)
}
