package service

import (
	"github.com/maxim12233/crypto-app-server/account/models"
	"github.com/maxim12233/crypto-app-server/account/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	GetAccount(id uint) (*models.Account, error)
	DeleteAccount(id uint) error
	CreateAccount(a models.CreateAccountRequest) error
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

func (s AccountService) GetAccount(id uint) (*models.Account, error) {
	account, err := s.repo.GetAccountById(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s AccountService) DeleteAccount(id uint) error {
	err := s.repo.DeleteAccountById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s AccountService) CreateAccount(a models.CreateAccountRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), 10)
	if err != nil {
		s.logger.Error("Error while hashing the password", zap.Error(err))
		return err
	}
	acc := models.Account{
		Login:        a.Login,
		Email:        a.Email,
		PasswordHash: string(hash),
	}
	err = s.repo.CreateAccount(acc)
	if err != nil {
		return err
	}
	return nil
}
