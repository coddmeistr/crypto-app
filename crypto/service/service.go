package service

import (
	"errors"

	cryptocompare "github.com/maxim12233/crypto-app-server/crypto/crypto_compare_sdk"
	"github.com/maxim12233/crypto-app-server/crypto/models"
	"github.com/maxim12233/crypto-app-server/crypto/repository"
	"go.uber.org/zap"
)

type ICryptoService interface {
	GetPrice(sym string, symsTo []string) (*cryptocompare.Prices, error)
	GetHistory(timebase string, fsym string, tsym string, limit int) (*cryptocompare.HistoricalData, error)
	GetPriceDifference(timebase string, symbol string, symbolTo string, offset int) (*models.PriceDifference, error)
}

type CryptoService struct {
	repo   repository.IAccountRepository
	market cryptocompare.ICryptoCompare
	logger *zap.Logger
}

func NewCryptoService(repo repository.IAccountRepository, logger *zap.Logger, market cryptocompare.ICryptoCompare) ICryptoService {
	return CryptoService{
		repo:   repo,
		logger: logger,
		market: market,
	}
}

func (s CryptoService) GetPriceDifference(timebase string, symbol string, symbolTo string, offset int) (*models.PriceDifference, error) {

	latest, err := s.GetPrice(symbol, []string{symbolTo})
	if err != nil {
		return nil, err
	}

	history, err := s.GetHistory(timebase, symbol, symbolTo, offset)
	if err != nil {
		return nil, err
	}

	if _, ok := latest.Prices[symbolTo]; !ok {
		return nil, errors.New("Failed to get latest price according to given symbolTo. Map doesn't have symbolTo key")
	}

	latestPrice := latest.Prices[symbolTo]
	historyPrice := (history.Data[0].High + history.Data[0].Low) / 2

	return &models.PriceDifference{
		Diff:         latestPrice - historyPrice,
		DiffPercents: ((latestPrice - historyPrice) / historyPrice) * 100,
	}, nil
}

func (s CryptoService) GetPrice(sym string, symsTo []string) (*cryptocompare.Prices, error) {

	prices, err := s.market.GetLatestPrice(sym, symsTo)
	if err != nil {
		return nil, err
	}
	return prices, nil
}

func (s CryptoService) GetHistory(timebase string, fsym string, tsym string, limit int) (*cryptocompare.HistoricalData, error) {

	var data *cryptocompare.HistoricalData
	var err error
	switch timebase {
	case "days":
		data, err = s.market.GetHistoricalDailyOHLCV(fsym, tsym, limit)
	case "hours":
		data, err = s.market.GetHistoricalHourlyOHLCV(fsym, tsym, limit)
	case "minutes":
		data, err = s.market.GetHistoricalMinutlyOHLCV(fsym, tsym, limit)
	default:
		return nil, errors.New("Invalid timebase param")
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
