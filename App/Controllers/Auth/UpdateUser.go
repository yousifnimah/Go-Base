package Auth

import (
	"github.com/gin-gonic/gin"
	"base/App/Handlers/Bcrypt"
	"base/App/Handlers/GORM"
	"base/App/Handlers/JWT"
	"base/App/Handlers/Redis"
	"base/App/Handlers/Validator"
	"base/App/Models"
	"base/Helper"
	"net/http"
	"time"
)

type UpdateRequest struct {
	NewPassword             string `form:"new_password" json:"new_password" xml:"new_password" validate:"omitempty,min=8"`
	NewPasswordConfirmation string `form:"new_password_confirmation" json:"new_password_confirmation" xml:"new_password_confirmation" validate:"required_if=NewPassword ALLOW,eqfield=NewPassword"`
	FullName                string `form:"full_name" json:"full_name" xml:"full_name" validate:"required,max=20"`
}

func UpdateUser(c *gin.Context) {
	var RequestBody UpdateRequest
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
	if err := db.First(&User, "id = ?", JWT.Claims.ID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": struct {
			Username []string `json:"username"`
		}{
			Username: []string{Helper.Localize(c, "invalid_username")},
		},
		})
		c.Abort()
		return
	}
	User.FullName = RequestBody.FullName
	if RequestBody.NewPassword != "" {
		User.Password = Bcrypt.HashPassword(RequestBody.NewPassword)
	}
	db.Save(&User)

	if RequestBody.NewPassword != "" {
		err = Redis.InvokeAllAccess(User.Username)
		if err != nil {
			return
		}

		expireHours := 8760
		expirationTime := time.Now().Add(time.Duration(expireHours) * time.Hour)
		jwt, err := JWT.GenerateJWT(&User, expirationTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
			c.Abort()
			return
		}
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
	} else {
		c.JSON(
			http.StatusOK,
			GetUserResp{
				User: UserClaim{
					ID:       User.ID,
					Username: User.Username,
					FullName: User.FullName,
				},
			},
		)
	}
	c.Abort()
	GORM.CloseConnection(db)
	return
}
