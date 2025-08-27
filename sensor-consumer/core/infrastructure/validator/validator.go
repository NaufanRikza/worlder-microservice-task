package validator

import go_validator "github.com/go-playground/validator/v10"

type validator struct {
	validator *go_validator.Validate
}

type Validator interface {
	Validate(i interface{}) error
}

func NewValidator() Validator {
	return &validator{validator: go_validator.New()}
}

func (cv *validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

