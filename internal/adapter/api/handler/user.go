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

// GetCurrentUser godoc
// @Summary      获取当前登录用户
// @Description  返回当前登录用户的详细信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Success      200  {object}  viewmodel.User
// @Failure      401  {object}  map[string]any
// @Router       /user [get]
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	// TODO: 从认证中间件中获取用户ID
	// userID := auth.GetUserID(c)

	// TODO: 使用用户服务获取用户信息
	// user, err := h.userService.GetUser(userID)
	// if err != nil {
	//    c.JSON(http.StatusUnauthorized, gin.H{"error": "need login"})
	//     return
	// }

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
