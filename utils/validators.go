package utils

import (
	"github.com/go-playground/validator"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

func ValidatorErrors(err error) string {
	var errorMessage string
	for _, err := range err.(validator.ValidationErrors) {
		errorMessage = "Error in field: " + err.Field()
	}
	return errorMessage

}
