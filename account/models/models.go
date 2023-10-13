package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
	Email        string `json:"email"`
}

type CreateAccountRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
