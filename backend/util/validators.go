package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateInputs(dataSet interface{}) map[string]string {
	validate = validator.New()
	err := validate.Struct(dataSet)

	if err != nil {
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("%s is a required field", err.Field())
		}

		return errors
	}

	return nil
}
