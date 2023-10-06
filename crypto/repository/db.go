package repository

import (
	"fmt"

	"github.com/maxim12233/crypto-app-server/crypto/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbUrl string) (*gorm.DB, error) {

	fmt.Println(dbUrl)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Currency{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
