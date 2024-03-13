package Routes

import (
	"github.com/gin-gonic/gin"
	"base/App/Controllers/Auth"
	"base/App/Middlewares"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		auth := v1.Group("auth")
		{
			loggedOut := auth.Group("").Use(Middlewares.Guest())
			{
				loggedOut.POST("login", Auth.Login)
			}
			loggedIn := auth.Group("").Use(Middlewares.Auth())
			{
				loggedIn.GET("user", Auth.GetUser)
				loggedIn.PATCH("update-user", Auth.UpdateUser)
				loggedIn.DELETE("logout", Auth.Logout)
			}
		}
	}
}
