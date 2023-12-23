package service

import (
	"errors"
	"fmt"
	"strings"

	app "github.com/maxim12233/crypto-app-server/account"
	"github.com/maxim12233/crypto-app-server/account/models"
	"github.com/maxim12233/crypto-app-server/account/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	GetAccount(id uint) (*models.Account, error)
	DeleteAccount(id uint) error
	CreateAccount(login string, password string, email string) error
	Login(login string, password string, email string) (uint, []uint, error)
	GetBalance(id uint) (*models.Balance, error)
	BuyActivity(id uint, symbol string, price float64) error
	SellActivity(id uint, symbol string, price float64, amount float64) error
	FakeDeposit(id uint, amount float64) error
	GetActivities(id uint, symbols string, fetchPrices bool) ([]models.Activity, map[string]float64, error)
	GetActivity(id uint, symbol string) (*models.Activity, error)
	GetAllAccounts() ([]models.Account, error)
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

func (s AccountService) GetAllAccounts() ([]models.Account, error) {
	accs, err := s.repo.GetAllAccounts()
	if err != nil {
		return nil, err
	}

	return accs, nil
}

func (s AccountService) GetActivities(id uint, symbols string, fetchPrices bool) ([]models.Activity, map[string]float64, error) {

	var symbolsSlice []string
	if symbols == "" {
		symbolsSlice = make([]string, 0)
	} else {
		symbolsSlice = strings.Split(symbols, ",")
	}
	activities, err := s.repo.GetActivities([]uint{id}, symbolsSlice)
	if err != nil {
		return nil, nil, err
	}

	var prices = make(map[string]float64)
	if fetchPrices {
		for _, v := range activities {
			price, err := fetchSymbolPriceFromCryptoMicroservice(v.Symbol, "USD", s.logger)
			if err != nil {
				return nil, nil, err
			}
			prices[v.Symbol] = price
		}
	}

	return activities, prices, nil
}

func (s AccountService) GetActivity(id uint, symbol string) (*models.Activity, error) {
	act, err := s.repo.GetActivity(id, symbol)
	if err != nil {
		return nil, err
	}
	return act, nil
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

func (s AccountService) SellActivity(id uint, symbol string, price float64, amount float64) error {

	if price < 0 || amount < 0 || (price == 0 && amount == 0) {
		s.logger.Error("Incorrect price or amount params < 0 or both == 0, should've been validated outside this function")
		return app.ErrBadRequest
	}

	// Balance
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

	currentSymbolPrice, err := fetchSymbolPriceFromCryptoMicroservice(symbol, "USD", s.logger)
	if err != nil {
		return err
	}

	// How much of crypto currency user sells(f.e. 1 BTC or 0.001 BTC)
	var haveUSD float64
	if price != 0 {
		haveUSD = hadUSD + price
		amount = price / currentSymbolPrice
	} else if amount != 0 {
		haveUSD = hadUSD + (amount * currentSymbolPrice)
	} else {
		s.logger.Error("Error, both price and amount are 0(zero)")
		return app.ErrBadRequest
	}

	// Change user's account balance and update it to BD
	balance.USD = &haveUSD
	if err := s.repo.UpdateAccountBalance(balance); err != nil {
		s.logger.Error("Error while updating user's balance", zap.Error(err))
		return err
	}

	// If user already has some of this cryptocurrency, then update current one
	// If after sell transaction user have 0 of this cryptocurrency amount then delete this record
	// If something goes wrong, then user's balance will be rollbacked with rollback function
	existingActivity, err := s.repo.GetActivity(id, symbol)
	if err != nil {
		if !errors.Is(err, app.ErrNotFound) {
			_ = rollbackBalance() // We're not handling rollback's error, because this function throws Fatal anyway
			s.logger.Error("Error getting existing activity", zap.Error(err))
			return err
		}
		_ = rollbackBalance()
		s.logger.Error("Error activity doesn't exist. Cannot sell not existing activity", zap.Error(err))
		return err
	}

	resultAmount := existingActivity.Amount - amount
	if resultAmount < 0 {

		_ = rollbackBalance()
		s.logger.Error("result amount < 0. User dont have enough cryptocurrency to sell it.")
		return app.ErrNotEnoughCurrency

	} else if resultAmount == 0 {

		err := s.repo.DeleteActivity(id, symbol)
		if err != nil {
			_ = rollbackBalance()
			s.logger.Error("Error deleting activity.", zap.Error(err))
			return err
		}

	} else {

		existingActivity.Amount = resultAmount
		if err := s.repo.UpdateActivity(existingActivity); err != nil {
			_ = rollbackBalance()
			s.logger.Error("Error updating activity", zap.Error(err))
			return err
		}

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

	currentSymbolPrice, err := fetchSymbolPriceFromCryptoMicroservice(symbol, "USD", s.logger)
	if err != nil {
		return err
	}

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

func (s AccountService) Login(login string, password string, email string) (uint, []uint, error) {

	var matchesAccount *models.Account
	var err error
	if login != "" {
		if matchesAccount, err = s.repo.GetAccountByLogin(login); err != nil {
			s.logger.Error("Error getting account by login", zap.Error(err))
			return 0, nil, err
		}
	} else if email != "" {
		if matchesAccount, err = s.repo.GetAccountByEmail(email); err != nil {
			s.logger.Error("Error getting account by email", zap.Error(err))
			return 0, nil, err
		}
	} else {
		s.logger.Error("Error login and email are empty strings", zap.Error(err))
		return 0, nil, app.WrapE(app.ErrBadRequest, "Empty login and email fields")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(matchesAccount.PasswordHash), []byte(password)); err != nil {
		s.logger.Error("Error passwords not match", zap.Error(err))
		return 0, nil, app.ErrIncorrectLoginOrPassword
	}

	roles := make([]uint, 0)
	for _, v := range matchesAccount.AccountRole {
		roles = append(roles, v.RoleID)
	}

	fmt.Println(matchesAccount)

	fmt.Println(roles)
	return matchesAccount.ID, roles, nil
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

	roles := make([]models.AccountRole, 1)
	roles = append(roles, models.AccountRole{
		RoleID: 1,
	})
	if login == "admin" {
		roles = append(roles, models.AccountRole{
			RoleID: 2,
		})
		roles = append(roles, models.AccountRole{
			RoleID: 3,
		})
	}

	acc := models.Account{
		Login:        login,
		Email:        email,
		PasswordHash: string(hash),
		Balance: models.Balance{
			USD: &USD,
		},
		AccountRole: roles,
	}

	err = s.repo.CreateAccount(acc)
	if err != nil {
		s.logger.Error("Error creating account", zap.Error(err))
		return err
	}
	return nil
}
