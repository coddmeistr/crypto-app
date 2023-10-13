package repository

import (
	"github.com/maxim12233/crypto-app-server/account/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountById(id uint) (*models.Account, error)
	CreateAccount(a models.Account) error
	DeleteAccountById(id uint) error
}

type AccountRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAccountRepository(db *gorm.DB, logger *zap.Logger) IAccountRepository {
	return &AccountRepository{
		db:     db,
		logger: logger,
	}
}

func (r *AccountRepository) GetAccountById(id uint) (*models.Account, error) {
	var a models.Account
	result := r.db.Find(&a, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &a, nil
}

func (r *AccountRepository) CreateAccount(a models.Account) error {
	result := r.db.Create(&a)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AccountRepository) DeleteAccountById(id uint) error {
	result := r.db.Model(&models.Account{}).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
