package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IAccountValidator interface {
	AccountValidate(account model.Account) error
}

type accountValidator struct{}

func NewAccountValidator() IAccountValidator {
	return &accountValidator{}
}

func (av *accountValidator) AccountValidate(account model.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(
			&account.Name,
			validation.RuneLength(0, 20).Error("limited name max 20 char "),
		),
		validation.Field(
			&account.Introduction,
			validation.RuneLength(0, 100).Error("limited introduction max 100 char"),
		),
	)
}
