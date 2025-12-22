package middleware

import (
	"github.com/gin-gonic/gin"
	"go_blog/utils"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "Authorization header is required")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			utils.Unauthorized(c, "Bearer header is required")
			c.Abort()
			return
		}
		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.Unauthorized(c, "token is Invalid")
			c.Abort()
			return
		}
		//将token用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
