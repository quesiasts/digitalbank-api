package models

import (
	"time"

	"gopkg.in/validator.v2"
)

type Account struct {
	Id         string    `json:"id"`
	Name       string    `json:"name,omitempty"`
	Cpf        string    `json:"cpf" validate:"len=11,regexp=^[0-9]*$"`
	Secret     string    `json:"secret" validate:"nonzero"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at,omitempty"`
}

func ValidateAccount(account *Account) error {
	if err := validator.Validate(account); err != nil {
		return err
	}
	return nil
}
