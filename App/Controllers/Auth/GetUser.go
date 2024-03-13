package Auth

import (
	"base/App/Handlers/JWT"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserResp struct {
	User UserClaim `json:"user"`
}

func GetUser(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		GetUserResp{
			User: UserClaim{
				ID:       JWT.Claims.ID,
				Username: JWT.Claims.Username,
				FullName: JWT.Claims.FullName,
			},
		},
	)
}
