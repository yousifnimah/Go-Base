package Middlewares

import (
	"base/App/Handlers/JWT"
	"base/Helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Next()
		}
		err := JWT.ValidateJWT(tokenString)
		if err != nil {
			c.Next()
		}
		if tokenString != "" && err == nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"errors": struct {
				Authorization []string `json:"authorization"`
			}{
				Authorization: []string{Helper.Localize(c, "token_not_required")},
			},
			})
			c.Abort()
		}
		return
	}
}
