package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	*validator.Validate
}

func New() *Validator {
	return &Validator{validator.New(validator.WithRequiredStructEnabled())}
}
