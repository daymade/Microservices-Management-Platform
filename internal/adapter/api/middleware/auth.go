package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth 简单的token验证逻辑
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}
		// 在实际应用中，这里应该验证token的有效性
		// TODO 验证token的有效性

		c.Next()
	}
}
