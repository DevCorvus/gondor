package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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

func (user User) ComparePassword(password string) bool {
	hashedPasswordBytes := []byte(user.Password)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)

	return err == nil
}
