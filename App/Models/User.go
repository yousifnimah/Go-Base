package Models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"column:id"`
	FullName  string    `gorm:"column:full_name"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password;"`
	CreatedAt time.Time `gorm:"column:created_at;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;<-:update"`
}

type UserClaim struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
