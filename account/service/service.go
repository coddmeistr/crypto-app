package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"

	app "github.com/maxim12233/crypto-app-server/account"
	"github.com/maxim12233/crypto-app-server/account/config"
	"github.com/maxim12233/crypto-app-server/account/models"
	"github.com/maxim12233/crypto-app-server/account/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	GetAccount(id uint) (*models.Account, error)
	DeleteAccount(id uint) error
	CreateAccount(login string, password string, email string) error
	Login(login string, password string, email string) (uint, error)
	GetBalance(id uint) (*models.Balance, error)
	BuyActivity(id uint, symbol string, price float64) error
	FakeDeposit(id uint, amount float64) error
	GetActivities(id uint, symbols string) ([]models.Activity, error)
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

func (s AccountService) GetActivities(id uint, symbols string) ([]models.Activity, error) {

	var symbolsSlice []string
	if symbols == "" {
		symbolsSlice = make([]string, 0)
	} else {
		symbolsSlice = strings.Split(symbols, ",")
	}
	activities, err := s.repo.GetActivities([]uint{id}, symbolsSlice)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (s AccountService) FakeDeposit(id uint, amount float64) error {
	balance, err := s.repo.GetAccountBalance(id)
	if err != nil {
		s.logger.Error("Error getting account balance", zap.Error(err))
		return err
	}
	if balance.USD == nil {
		s.logger.Error("Balance.USD is nil. Critical error")
		return app.ErrInternal
	}
	sum := *balance.USD + amount
	balance.USD = &sum
	if err := s.repo.UpdateAccountBalance(balance); err != nil {
		s.logger.Error("Error when updating account balance", zap.Error(err))
		return app.ErrInternal
	}
	return nil
}

func (s AccountService) BuyActivity(id uint, symbol string, price float64) error {

	balance, err := s.repo.GetAccountBalance(id)
	if err != nil {
		s.logger.Error("Error getting account balance", zap.Error(err))
		return err
	}
	if balance.USD == nil {
		s.logger.Error("Error in balance structure, USD field is NIL, must be at least pointer to 0")
		return app.ErrInternal
	}
	hadUSD := *balance.USD
	haveUSD := hadUSD - price
	if haveUSD < 0 {
		s.logger.Error("Error not enough account balance", zap.Error(err))
		return app.ErrNotEnoughBalance
	}

	// Function which calls to rollback user balance to the start value if something goes wrong in activity operations
	// It calls fatal logger action, to prevent further errors due to failing rollback balance operation
	rollbackBalance := func() error {
		balance.USD = &hadUSD
		if err := s.repo.UpdateAccountBalance(balance); err != nil {
			s.logger.Fatal("FATAL Error while rollbacking(updating) user's balance. Some of the batabase data now corrupted", zap.Error(err))
			return err
		}
		return nil
	}

	// Take the current price for symbol via asking crypto microservice
	// Using config values to get host and path to foreign microservice
	var currentSymbolPrice float64
	host := config.GetConfig().GetString("dependencies.crypto_service.host")
	resPath := config.GetConfig().GetString("dependencies.crypto_service.endpoints.current_prices")
	uri, err := url.ParseRequestURI(host)
	if err != nil {
		s.logger.Error("Couldn't parse uri from the config host string", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	uri.Path = path.Join(uri.Path, resPath)
	q := uri.Query()
	q.Set("symbol", symbol)
	q.Set("symbolsTo", "USD")
	uri.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		s.logger.Error("Couldn't create new request to request crypto service", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Error("Error while doing request to foreign api", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	if resp.StatusCode != http.StatusOK {
		s.logger.Error("Foreign api, got not 200 code response")
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	var body models.Response // Basic response model for all of this app's microservices and gateway
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		s.logger.Error("Error while decoding json from crypto service response", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	// Get the payload response struct
	var payload interface{}
	if body.Payload == nil {
		s.logger.Error("Error, got an empty nil payload from crypto service", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	payload = *body.Payload
	pricesMap, ok := payload.(map[string]interface{})
	if !ok {
		s.logger.Error("Cannot cast payload to proper response type", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	prices, ok := pricesMap["Prices"].(map[string]interface{})
	if !ok {
		s.logger.Error("Cannot cast payload to proper response type", zap.Error(err))
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	// Check if needed value exists
	if _, ok := prices["USD"]; !ok {
		s.logger.Error("Final map dont contain USD key")
		return app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	// Get this value and try to covert to wanted type
	currentSymbolPrice = prices["USD"].(float64)

	// How much of crypto currency user buys(f.e. 1 BTC or 0.001 BTC)
	userBuys := price / currentSymbolPrice

	// Change user's account balance and update it to BD
	balance.USD = &haveUSD
	if err := s.repo.UpdateAccountBalance(balance); err != nil {
		s.logger.Error("Error while updating user's balance", zap.Error(err))
		return err
	}

	// If user already has some of this cryptocurrency, then update current one
	// If user dont have it, then create new activity for this user
	// If something goes wrong, then user's balance will be rollbacked with rollback function
	existingActivity, err := s.repo.GetActivity(id, symbol)
	if err != nil {
		if !errors.Is(err, app.ErrNotFound) {
			_ = rollbackBalance() // We're not handling rollback's error, because this function throws Fatal anyway
			s.logger.Error("Error creating activity", zap.Error(err))
			return err
		}

		newActivity := &models.Activity{
			AccountID: id,
			Symbol:    symbol,
			Amount:    userBuys,
		}
		if err := s.repo.CreateActivity(newActivity); err != nil {
			_ = rollbackBalance()
			s.logger.Error("Error creating activity", zap.Error(err))
			return err
		}
	} else {
		existingActivity.Amount += userBuys
		if err := s.repo.UpdateActivity(existingActivity); err != nil {
			_ = rollbackBalance()
			s.logger.Error("Error updating activity", zap.Error(err))
			return err
		}
	}

	return nil
}

func (s AccountService) GetBalance(accid uint) (*models.Balance, error) {

	balance, err := s.repo.GetAccountBalance(accid)
	if err != nil {
		s.logger.Error("Error getting account balance", zap.Error(err))
		return nil, err
	}
	return balance, nil
}

func (s AccountService) Login(login string, password string, email string) (uint, error) {

	var matchesAccount *models.Account
	var err error
	if login != "" {
		if matchesAccount, err = s.repo.GetAccountByLogin(login); err != nil {
			s.logger.Error("Error getting account by login", zap.Error(err))
			return 0, err
		}
	} else if email != "" {
		if matchesAccount, err = s.repo.GetAccountByEmail(email); err != nil {
			s.logger.Error("Error getting account by email", zap.Error(err))
			return 0, err
		}
	} else {
		s.logger.Error("Error login and email are empty strings", zap.Error(err))
		return 0, app.WrapE(app.ErrBadRequest, "Empty login and email fields")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(matchesAccount.PasswordHash), []byte(password)); err != nil {
		s.logger.Error("Error passwords not match", zap.Error(err))
		return 0, app.ErrIncorrectLoginOrPassword
	}

	return matchesAccount.ID, nil
}

func (s AccountService) GetAccount(id uint) (*models.Account, error) {
	account, err := s.repo.GetAccountById(id)
	if err != nil {
		s.logger.Error("Error getting account by id", zap.Error(err))
		return nil, err
	}
	return account, nil
}

func (s AccountService) DeleteAccount(id uint) error {
	err := s.repo.DeleteAccountById(id)
	if err != nil {
		s.logger.Error("Error deleting account", zap.Error(err))
		return err
	}
	return nil
}

func (s AccountService) CreateAccount(login string, password string, email string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		s.logger.Error("Error while hashing the password", zap.Error(err))
		return app.ErrInternal
	}

	var (
		USD = float64(15000)
	)
	acc := models.Account{
		Login:        login,
		Email:        email,
		PasswordHash: string(hash),
		Balance: models.Balance{
			USD: &USD,
		},
	}

	err = s.repo.CreateAccount(acc)
	if err != nil {
		s.logger.Error("Error creating account", zap.Error(err))
		return err
	}
	return nil
}
