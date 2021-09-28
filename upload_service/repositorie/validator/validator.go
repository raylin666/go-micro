package validator

import (
	"github.com/raylin666/go-micro-protoc/errors"
	"github.com/raylin666/go-utils/validator"
)

var (
	instance *validator.Validator
)

func New(locale string) *validator.Validator {
	instance = validator.New(locale)
	return instance
}

func Validate(data interface{}) error {
	reason := instance.Validate(data)
	if len(reason) > 0 {
		return errors.ErrorDataValidateError(reason)
	}

	return nil
}

