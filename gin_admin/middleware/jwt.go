package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("my-secret-2026")

// JWT 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "未登录"})
			c.Abort()
			return
		}

		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "令牌无效"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Web 页面 Cookie 登录验证
func JWTAuthCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, _ := c.Cookie("token")
		if tokenStr == "" {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
