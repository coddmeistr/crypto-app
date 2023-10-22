package repository

import (
	"errors"

	app "github.com/maxim12233/crypto-app-server/account"
	"github.com/maxim12233/crypto-app-server/account/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountById(id uint) (*models.Account, error)
	GetAccountByLogin(login string) (*models.Account, error)
	GetAccountByEmail(email string) (*models.Account, error)
	CreateAccount(a models.Account) error
	DeleteAccountById(id uint) error
	GetAccountBalance(accid uint) (*models.Balance, error)
	UpdateAccountBalance(b *models.Balance) error
	UpdateActivity(act *models.Activity) error
	GetActivity(accid uint, symbol string) (*models.Activity, error)
	CreateActivity(act *models.Activity) error
	GetActivities(accids []uint, symbols []string) ([]models.Activity, error)
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

func (r *AccountRepository) GetActivities(accids []uint, symbols []string) ([]models.Activity, error) {
	var activities []models.Activity
	var result *gorm.DB
	if len(accids) == 0 && len(symbols) == 0 {
		result = r.db.Find(&activities)
	} else if len(accids) == 0 {
		result = r.db.Where("symbo IN ?", symbols).Find(&activities)
	} else if len(symbols) == 0 {
		result = r.db.Where("account_id IN ?", accids).Find(&activities)
	} else {
		result = r.db.Where("(account_id IN ?) AND (symbol IN ?)", accids, symbols).Find(&activities)
	}
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			r.logger.Error("Get activities not found", zap.Error(result.Error))
			return nil, app.ErrNotFound
		}
		r.logger.Error("Get activities error", zap.Error(result.Error))
		return nil, app.ErrInternal
	}

	return activities, nil
}

func (r *AccountRepository) UpdateActivity(act *models.Activity) error {

	result := r.db.Model(&act).Updates(&models.Activity{Amount: act.Amount})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Update activity error", zap.Error(result.Error))
			return app.ErrNotFound
		}
		r.logger.Error("Update activity error", zap.Error(result.Error))
		return app.ErrInternal
	}

	return nil
}

func (r *AccountRepository) GetActivity(accid uint, symbol string) (*models.Activity, error) {

	var act models.Activity
	result := r.db.Where("account_id = ? AND symbol = ?", accid, symbol).First(&act)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Get activity error", zap.Error(result.Error))
			return nil, app.ErrNotFound
		}
		r.logger.Error("Get activity error", zap.Error(result.Error))
		return nil, app.ErrInternal
	}
	return &act, nil
}

func (r *AccountRepository) CreateActivity(act *models.Activity) error {
	result := r.db.Create(&act)
	if result.Error != nil {
		r.logger.Error("Create activity error", zap.Error(result.Error))
		return app.ErrInternal
	}

	return nil
}

func (r *AccountRepository) UpdateAccountBalance(b *models.Balance) error {

	result := r.db.Save(&b)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Update balance error", zap.Error(result.Error))
			return app.ErrNotFound
		}
		r.logger.Error("Update balance error", zap.Error(result.Error))
		return app.ErrInternal
	}

	return nil
}

func (r *AccountRepository) GetAccountBalance(accid uint) (*models.Balance, error) {

	var b models.Balance
	result := r.db.Where("account_id = ?", accid).First(&b)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Get balance error", zap.Error(result.Error))
			return nil, app.ErrNotFound
		}
		r.logger.Error("Get balance error", zap.Error(result.Error))
		return nil, app.ErrInternal
	}
	return &b, nil
}

func (r *AccountRepository) GetAccountById(id uint) (*models.Account, error) {

	var a models.Account
	result := r.db.First(&a, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Get account by id error", zap.Error(result.Error))
			return nil, app.ErrNotFound
		}
		r.logger.Error("Get account by id error", zap.Error(result.Error))
		return nil, app.ErrInternal
	}
	return &a, nil
}

func (r *AccountRepository) GetAccountByEmail(email string) (*models.Account, error) {

	var a models.Account
	result := r.db.Where("email = ?", email).First(&a)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Get account by email error", zap.Error(result.Error))
			return nil, app.ErrNotFound
		}
		r.logger.Error("Get account by email error", zap.Error(result.Error))
		return nil, app.ErrInternal
	}
	return &a, nil
}

func (r *AccountRepository) GetAccountByLogin(login string) (*models.Account, error) {

	var a models.Account
	result := r.db.Where("login = ?", login).First(&a)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Error("Get account by login error", zap.Error(result.Error))
			return nil, app.ErrNotFound
		}
		r.logger.Error("Get account by login error", zap.Error(result.Error))
		return nil, app.ErrInternal
	}
	return &a, nil
}

func (r *AccountRepository) CreateAccount(a models.Account) error {
	result := r.db.Create(&a)
	if result.Error != nil {
		r.logger.Error("Create account error", zap.Error(result.Error))
		return app.ErrInternal
	}

	return nil
}

func (r *AccountRepository) DeleteAccountById(id uint) error {
	result := r.db.Model(&models.Account{}).Delete(id)
	if result.Error != nil {
		r.logger.Error("Delete account error", zap.Error(result.Error))
		return app.ErrInternal
	}
	return nil
}
