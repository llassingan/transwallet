package app

import "github.com/go-playground/validator/v10"

func minAccountID(fl validator.FieldLevel) bool {
	return fl.Field().Int() >= 100001
}

func NewValidator(options []validator.Option) *validator.Validate{
	validate := validator.New(options...)
	validate.RegisterValidation("minaccountid", minAccountID)
	return validate
}