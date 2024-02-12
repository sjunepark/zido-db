package validation

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func Init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(i interface{}) error {
	return validate.Struct(i)
}
