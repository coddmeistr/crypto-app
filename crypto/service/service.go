package service

import (
	coinmarket "github.com/maxim12233/crypto-app-server/crypto/coinmarketsdk"
	"github.com/maxim12233/crypto-app-server/crypto/repository"
	"go.uber.org/zap"
)

type ICryptoService interface {
	GetQuote(slug string) (coinmarket.USD, error)
}

type CryptoService struct {
	repo   repository.IAccountRepository
	market coinmarket.ICoinMarket
	logger *zap.Logger
}

func NewCryptoService(repo repository.IAccountRepository, logger *zap.Logger, market coinmarket.ICoinMarket) ICryptoService {
	return CryptoService{
		repo:   repo,
		logger: logger,
		market: market,
	}
}

func (s CryptoService) GetQuote(slug string) (coinmarket.USD, error) {

	quote, err := s.market.GetLatestQuotes(slug)
	if err != nil {
		return coinmarket.USD{}, err
	}
	return quote, nil
}
