package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Name      string          `json:"name" validate:"required,min=2,max=100"`
	Email     string          `json:"email" validate:"required,email,max=250"`
	Password  string          `json:"password" validate:"required,min=6,max=500"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
