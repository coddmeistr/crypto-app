package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
	Email        string `json:"email"`
}
