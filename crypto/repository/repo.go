package repository

import (
	"github.com/maxim12233/crypto-app-server/crypto/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountInfoById(id uint) (*models.Currency, error)
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

func (r *AccountRepository) GetAccountInfoById(id uint) (*models.Currency, error) {
	return &models.Currency{
		Name:  "USD",
		ToUSD: 1,
	}, nil
}
