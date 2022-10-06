package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Email     string          `json:"email" validate:"required,email"`
	Password  string          `json:"password" validate:"required,min=6"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
