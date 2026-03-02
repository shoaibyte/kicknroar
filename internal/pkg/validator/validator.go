package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	
	// Register custom tag name function to use json tags
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate validates a struct
func Validate(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		var errorMessages []string
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, getErrorMessage(validationErr))
		}
		return fmt.Errorf("%s", strings.Join(errorMessages, "; "))
	}
	return nil
}

// getErrorMessage returns a user-friendly error message
func getErrorMessage(err validator.FieldError) string {
	field := err.Field()
	if field == "" {
		field = err.StructField()
	}

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, err.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, err.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

