package Auth

import (
	"github.com/gin-gonic/gin"
	"base/App/Handlers/Bcrypt"
	"base/App/Handlers/GORM"
	"base/App/Handlers/JWT"
	"base/App/Handlers/Validator"
	"base/App/Models"
	"base/Helper"
	"net/http"
	"time"
)

type Request struct {
	Username   string `form:"username" json:"username" xml:"username" validate:"required,exists=users"`
	Password   string `form:"password" json:"password" xml:"password" validate:"required"`
	RememberMe bool   `form:"remember_me" json:"remember_me" xml:"remember_me" validate:"omitempty"`
}

type Access struct {
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type UserClaim Models.UserClaim

type Response struct {
	User   UserClaim `json:"user"`
	Access Access    `json:"access"`
}

func Login(c *gin.Context) {
	var RequestBody Request
	err := c.Bind(&RequestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		c.Abort()
		return
	}
	Validator.NewValidator(c)
	if err := Validator.Validate(Validator.Validator.Struct(RequestBody)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err})
		c.Abort()
		return
	}

	var User Models.User
	db := GORM.OpenConnection()
	if err := db.First(&User, "username = ?", RequestBody.Username).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": struct {
			Username []string `json:"username"`
		}{
			Username: []string{Helper.Localize(c, "invalid_username")},
		},
		})
		c.Abort()
		return
	}
	if err := Bcrypt.CheckPassword(RequestBody.Password, User); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": struct {
			Password []string `json:"password"`
		}{
			Password: []string{Helper.Localize(c, "wrong_password")},
		},
		})
		c.Abort()
		return
	}

	expireHours := 2190
	if RequestBody.RememberMe {
		expireHours = 8760
	}
	expirationTime := time.Now().Add(time.Duration(expireHours) * time.Hour)
	jwt, err := JWT.GenerateJWT(&User, expirationTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		c.Abort()
		return
	}

	GORM.CloseConnection(db)

	c.JSON(
		http.StatusOK,
		Response{
			User: UserClaim{
				ID:       User.ID,
				Username: User.Username,
				FullName: User.FullName,
			},
			Access: Access{
				Token:          jwt,
				ExpirationDate: expirationTime,
			},
		},
	)
	c.Abort()
	return
}
