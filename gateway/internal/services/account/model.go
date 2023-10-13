package account_service

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Login string `json:"login"`
	Email string `json:"email"`
}
