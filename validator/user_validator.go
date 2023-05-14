package validator

import (
	"backend/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type UserValidator struct{}

func NewUserValidator() IUserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("require email"),
			validation.RuneLength(1, 30).Error("limit email"),
			is.Email.Error("correct email"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("require password"),
			validation.RuneLength(6, 30).Error("limit password"),
		),
	)
}