package utils

import (
	"github.com/go-playground/validator/v10"
)

// Create a new validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct based on the tags defined
func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		return errors
	}
	return nil
}
