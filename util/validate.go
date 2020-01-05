package util

import (
	"errors"
	"github.com/go-playground/validator"
	"strings"
)

func Validate(model interface{}) error {
	v := validator.New()
	err := v.Struct(model)
	if err != nil {
		var requiredField []string
		var inValidField []string
		var requiredText string
		var invalidText string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				requiredField = append(requiredField, err.StructField())
			case "email":
				inValidField = append(inValidField, err.StructField())
			}
		}

		if len(requiredField) < 2 && len(requiredField) > 0 {
			requiredText = strings.Join(requiredField, ", ") + " is required"
		} else if len(requiredField) >= 2 {
			requiredText = strings.Join(requiredField, ", ") + " are required"
		}

		if len(requiredField) > 0 && len(inValidField) > 0 {
			requiredText += "; "
		}

		if len(inValidField) < 2 && len(inValidField) > 0 {
			invalidText = strings.Join(inValidField, ", ") + " is invalid"
		} else if len(inValidField) >= 2  {
			invalidText =  strings.Join(inValidField, ", ") + " are invalid"
		}


		return errors.New(requiredText + invalidText)
	}
	return nil
}
