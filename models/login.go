package models

import (
	"gopkg.in/validator.v2"
)

type Login struct {
	Cpf    string `json:"cpf" validate:"len=11,regexp=^[0-9]*$"`
	Secret string `json:"secret" validate:"nonzero"`
}

func ValidateLogin(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}
	return nil
}
