package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(request interface{}) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}
func InitValidators() {
	validate = validator.New()
	_ = validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		compile := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		return compile.MatchString(fl.Field().String())
	})

	_ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		compile := regexp.MustCompile("^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[*.!@$%^&(){}[]:;<>,.?/~_+-=|\\]).{8,32}$")
		return compile.MatchString(fl.Field().String())
	})

}
