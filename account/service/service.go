package service

import (
	"github.com/maxim12233/crypto-app-server/account/models"
	"github.com/maxim12233/crypto-app-server/account/repository"
	"go.uber.org/zap"
)

type IAccountService interface {
	GetAccountInfoById(id uint) (*models.Account, error)
}

type AccountService struct {
	repo   repository.IAccountRepository
	logger *zap.Logger
}

func NewAccountService(repo repository.IAccountRepository, logger *zap.Logger) IAccountService {
	return AccountService{
		repo:   repo,
		logger: logger,
	}
}

func (s AccountService) GetAccountInfoById(id uint) (*models.Account, error) {
	account, err := s.repo.GetAccountInfoById(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
