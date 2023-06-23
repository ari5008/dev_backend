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
			&account.UserName,
			validation.RuneLength(0, 15).Error("limited user_name max 15 char "),
		),
	)
}
