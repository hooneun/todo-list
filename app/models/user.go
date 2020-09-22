package models

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName user model
func (User) TableName() string {
	return "users"
}
