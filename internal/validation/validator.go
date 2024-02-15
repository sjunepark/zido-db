package validation

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(i interface{}) error {
	return validate.Struct(i)
}
