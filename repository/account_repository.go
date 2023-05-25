package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAccountRepository interface {
	CreateAccount(account *model.Account) error
	GetAccount(account *model.Account, userId uint) error
	UpdateAccount(account *model.Account, userId uint, accountId uint) error
	DeleteAccount(userId uint, accountId uint) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &accountRepository{db}
}

func (ar *accountRepository) CreateAccount(account *model.Account) error {
	if err := ar.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (ar *accountRepository) GetAccount(account *model.Account, userId uint) error {
	if err := ar.db.Where("user_id=?", userId).First(account).Error; err != nil {
		return err
	}
	return nil
}

func (ar *accountRepository) UpdateAccount(account *model.Account, userId uint, accountId uint) error {
	result := ar.db.Model(account).Clauses(clause.Returning{}).Where("id=? AND user_id=?", accountId, userId).
		Updates(map[string]interface{}{
			"name":         account.Name,
			"image_url":    account.ImageURL,
			"introduction": account.Introduction,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (ar *accountRepository) DeleteAccount(userId uint, accountId uint) error {
	result := ar.db.Where("id=? AND user_id=?", accountId, userId).Delete(&model.Account{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
