package repository

import (
	"github.com/maxim12233/crypto-app-server/account/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbUrl string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Account{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
