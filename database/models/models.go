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
	Gophers   []Gopher        `json:"gophers"`
}

func (user User) ComparePassword(password string) bool {
	hashedPasswordBytes := []byte(user.Password)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)

	return err == nil
}

type Gopher struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Name      string          `json:"name" validate:"required,min=2,max=100"`
	Color     string          `json:"color" validate:"required,hexcolor"`
	Age       uint8           `json:"age" validate:"required,gt=0,lte=7"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	UserID    uint            `json:"userId"`
}
