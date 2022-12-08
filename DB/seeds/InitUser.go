package seeds

import (
	"gateway_api/App/Handlers/Bcrypt"
	"gateway_api/App/Models"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) error {
	user := Models.User{FullName: "Admin", Username: "admin", Password: Bcrypt.HashPassword("123456")}
	return db.Create(&user).Error
}
