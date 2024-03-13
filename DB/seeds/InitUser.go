package seeds

import (
	"gorm.io/gorm"
	"base/App/Handlers/Bcrypt"
	"base/App/Models"
)

func InitUser(db *gorm.DB) error {
	user := Models.User{FullName: "Admin", Username: "admin", Password: Bcrypt.HashPassword("123456")}
	return db.Create(&user).Error
}
