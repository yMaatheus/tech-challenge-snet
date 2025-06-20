package util

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

func ParseValidationError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			errors[ve.Field()] = ve.Tag()
		}
	}
	return errors
}
