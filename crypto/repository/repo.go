package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ICryptoRepository interface {
}

type AccountRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAccountRepository(db *gorm.DB, logger *zap.Logger) ICryptoRepository {
	return &AccountRepository{
		db:     db,
		logger: logger,
	}
}
