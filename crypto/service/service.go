package service

import (
	app "github.com/maxim12233/crypto-app-server/crypto"
	bitfinex "github.com/maxim12233/crypto-app-server/crypto/bitfinex_sdk"
	cryptocompare "github.com/maxim12233/crypto-app-server/crypto/crypto_compare_sdk"
	"github.com/maxim12233/crypto-app-server/crypto/models"
	"go.uber.org/zap"
)

type ICryptoService interface {
	GetPrice(sym string, symsTo []string) (*cryptocompare.Prices, error)
	GetHistory(timebase string, fsym string, tsym string, limit int) (*cryptocompare.HistoricalData, error)
	GetPriceDifference(timebase string, symbol string, symbolTo string, offset int) (*models.PriceDifference, error)
}

type CryptoService struct {
	market   cryptocompare.ICryptoCompare
	wsmarket bitfinex.IBitfinex
	logger   *zap.Logger
}

func NewCryptoService(logger *zap.Logger, market cryptocompare.ICryptoCompare, wsmarket bitfinex.IBitfinex) ICryptoService {
	return CryptoService{
		logger:   logger,
		market:   market,
		wsmarket: wsmarket,
	}
}

func (s CryptoService) GetPriceDifference(timebase string, symbol string, symbolTo string, offset int) (*models.PriceDifference, error) {

	latest, err := s.GetPrice(symbol, []string{symbolTo})
	if err != nil {
		s.logger.Error("Error while getting prices from foreign api", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Foreign API Error")
	}

	history, err := s.GetHistory(timebase, symbol, symbolTo, offset)
	if err != nil {
		s.logger.Error("Error while getting history from foreign api", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Foreign API Error")
	}

	if _, ok := latest.Prices[symbolTo]; !ok {
		s.logger.Error("Foreign api response doesn't have required map key")
		return nil, app.WrapE(app.ErrInternal, "Foreign API Error")
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
		s.logger.Error("Error while getting price from foreign api", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Foreign API Error")
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
		s.logger.Error("Given invalid timebase parameter")
		return nil, app.WrapE(app.ErrBadRequest, "Invalid timebase for history")
	}
	if err != nil {
		s.logger.Error("Error while getting history from foreign api", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Foreign API Error")
	}

	return data, nil
}
