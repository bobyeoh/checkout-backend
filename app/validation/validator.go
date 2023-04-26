package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// CustomValidator godoc
type CustomValidator struct {
	Validator *validator.Validate
}

// Init godoc
func Init() *CustomValidator {
	validate := validator.New()
	validate.RegisterValidation("time", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02 15:04:05", fl.Field().String())
		return err == nil
	})
	return &CustomValidator{Validator: validate}
}

// Validate godoc
func (v *CustomValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

// ProcessError godocs
func ProcessError(err error) validator.FieldError {
	return err.(validator.ValidationErrors)[0]
}
