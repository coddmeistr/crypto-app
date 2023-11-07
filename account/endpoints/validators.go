package endpoints

import (
	"github.com/go-playground/validator/v10"
)

func makePasswordValidator(minLength int, maxLength int) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {

		l := fl.Field().Len()
		if l < minLength || l > maxLength {
			return false
		}

		return true
	}
}

func makeLoginValidator(minLength int, maxLength int) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {

		l := fl.Field().Len()
		if l < minLength || l > maxLength {
			return false
		}

		return true
	}
}
