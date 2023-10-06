package repository

import (
	"github.com/maxim12233/crypto-app-server/account/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountInfoById(id uint) (*models.Account, error)
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

func (r *AccountRepository) GetAccountInfoById(id uint) (*models.Account, error) {
	return &models.Account{
		Login:        "login",
		PasswordHash: "rfedg",
		Email:        "euseew.maxim2015@yandex.ru",
	}, nil
}
