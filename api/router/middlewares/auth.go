package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sunho/engbreaker/api/auth"
	"github.com/sunho/engbreaker/api/model"
)

func User(c *gin.Context) model.User {
	return c.MustGet("user").(model.User) //ㅁㄴㅇㄹ
}

func parseToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if token != "" {
		return token
	}

	cookie, err := c.Cookie("token")
	if err == nil {
		return cookie
	}

	return ""
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := parseToken(c)
		if token == "" {
			c.AbortWithStatus(401)
			return
		}

		user, err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
