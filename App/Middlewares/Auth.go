package Middlewares

import (
	"github.com/gin-gonic/gin"
	"base/App/Handlers/JWT"
	"base/App/Handlers/Redis"
	"base/Helper"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": struct {
				Authorization []string `json:"authorization"`
			}{
				Authorization: []string{Helper.Localize(c, "token_required")},
			},
			})
			c.Abort()
			return
		}
		err := JWT.ValidateJWT(tokenString)
		errRedis := Redis.GetAccess(JWT.Claims.Username, tokenString)
		if err != nil || errRedis != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": struct {
				Authorization []string `json:"authorization"`
			}{
				Authorization: []string{Helper.Localize(c, "invalid_token")},
			},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
