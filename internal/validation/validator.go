package validation

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())
	}
}

func ValidateStruct(i interface{}) error {
	InitValidator()
	return validate.Struct(i)
}
