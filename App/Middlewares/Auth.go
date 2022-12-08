package Middlewares

import (
	"gateway_api/App/Handlers/JWT"
	"gateway_api/App/Handlers/Redis"
	"gateway_api/Helper"
	"github.com/gin-gonic/gin"
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
