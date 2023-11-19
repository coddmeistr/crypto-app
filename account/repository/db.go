package repository

import (
	"time"

	"github.com/maxim12233/crypto-app-server/account/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Hard coded gouruting block to wait until postgres container starts
// Oddly depends_on docker-compose param doesn't work
// In all test cases 5 seconds were enough for postgres to run
const HARD_WAIT_FOR_POSTGRES_CONTAINER_START = 5 // in seconds

func InitDB(dbUrl string) (*gorm.DB, error) {
	time.Sleep(time.Second * HARD_WAIT_FOR_POSTGRES_CONTAINER_START)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Account{}, &models.Balance{}, &models.Activity{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
