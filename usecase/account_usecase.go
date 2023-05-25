package usecase

import (
	"backend/model"
	"backend/repository"
)

type IAccountUsecase interface {
	CreateAccount(account model.Account)  error
	GetAccount(userId uint) (model.AccountResponse, error)
	UpdateAccount(account model.Account, userId uint, accountId uint) (model.AccountResponse, error)
	DeleteAccount(userId uint, accountId uint) error
}

type accountUsecase struct {
	ar repository.IAccountRepository
}

func NewAccountUsecase(ar repository.IAccountRepository) IAccountUsecase {
	return &accountUsecase{ar}
}

func (au *accountUsecase) CreateAccount(account model.Account) error {
	if err := au.ar.CreateAccount(&account); err != nil {
		return err
	}
	return  nil
}

func (au *accountUsecase) GetAccount(userId uint) (model.AccountResponse, error) {
	account := model.Account{}
	if err := au.ar.GetAccount(&account, userId); err != nil {
		return model.AccountResponse{}, err
	}
	resAccount := model.AccountResponse{
		ID:           account.ID,
		Name:         account.Name,
		ImageURL:     account.ImageURL,
		Introduction: account.Introduction,
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    account.UpdatedAt,
	}
	return resAccount, nil
}

func (au *accountUsecase) UpdateAccount(account model.Account, userId uint, accountId uint) (model.AccountResponse, error) {
	if err := au.ar.UpdateAccount(&account, userId, accountId); err != nil {
		return model.AccountResponse{}, err
	}
	resAccount := model.AccountResponse{
		ID:           account.ID,
		Name:         account.Name,
		ImageURL:     account.ImageURL,
		Introduction: account.Introduction,
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    account.UpdatedAt,
	}
	return resAccount, nil
}

func (au *accountUsecase) DeleteAccount(userId uint, accountId uint) error {
	if err := au.ar.DeleteAccount(userId, accountId); err != nil {
		return err
	}
	return nil
}
