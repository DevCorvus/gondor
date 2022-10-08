package utils

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type errorResponse struct {
	Field    string
	Required string
	Value    string
}

var validate = validator.New()

func ValidateStruct[T any](s T) []*errorResponse {
	var errors []*errorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.Field = err.StructNamespace()
			element.Required = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}
