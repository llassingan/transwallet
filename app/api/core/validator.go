package core

import "github.com/go-playground/validator/v10"


// create new validator for account mumber
func minAccountID(fl validator.FieldLevel) bool {
	return (fl.Field().Int() >= 100001 && fl.Field().Int() <= 999999)
}

func NewValidator(options []validator.Option) *validator.Validate{
	// create new validator 
	validate := validator.New(options...)
	// register custom validator
	validate.RegisterValidation("minaccountid", minAccountID)
	return validate
}