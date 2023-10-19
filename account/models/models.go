package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Login        string `gorm:"unique"`
	PasswordHash string
	Email        string `gorm:"unique"`
	Balance      Balance
	Activity     []Activity
}

type Balance struct {
	gorm.Model
	AccountID uint `gorm:"unique"`
	USD       *float64
}

type Activity struct {
	gorm.Model
	AccountID uint
	Symbol    string  `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
}

type BuyActivityRequest struct {
	Symbol string  `validate:"required"`
	Price  float64 `validate:"required,gt=0"`
}

type GetAccountBalanceResponse struct {
	AccountID uint    `json:"account_id"`
	USD       float64 `json:"usd"`
}

type CreateAccountRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required_without=Login,omitempty,email"`
}

type LoginResponse struct {
	UserID uint `json:"id"`
}

type GetAccountResponse struct {
	ID    uint   `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}
