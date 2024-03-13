package Auth

import (
	"github.com/gin-gonic/gin"
	"base/App/Handlers/JWT"
	"base/App/Handlers/Redis"
	"net/http"
)

func Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	_ = Redis.InvokeAccess(JWT.Claims.Username, tokenString)
	c.JSON(http.StatusOK, nil)
	c.Abort()
	return
}
