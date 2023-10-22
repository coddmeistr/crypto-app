package models

import (
	"gorm.io/gorm"
)

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
