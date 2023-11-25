package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/maxim12233/crypto-app-server/account/config"
	"github.com/maxim12233/crypto-app-server/account/endpoints"
	"github.com/maxim12233/crypto-app-server/account/repository"
	"github.com/maxim12233/crypto-app-server/account/service"
	"github.com/maxim12233/crypto-app-server/account/transport"
)

// @title Account Service API
// @version 1.0
// @description Swagger API for Golang Project Crypto Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email euseew.maxim2015@yandex.ru

// @license.name EUS

// @BasePath /v1/account
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Couldn't load env vars from .env file")
	}

	_ = flag.Bool("docker", false, "Defines if app runs with docker")
	flag.Parse()

	if err := config.Init("local"); err != nil {
		panic(err)
	}

	var dbUrl string
	dbUrl = os.Getenv("DB_CONNECTION_STRING")
	if dbUrl == "" {
		panic("DB_CONNECTION_STRING env is empty string")
	}
	dbSession, err := repository.InitDB(dbUrl)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}

	logger := config.InitializeLogger()

	repo := repository.NewAccountRepository(dbSession, logger)
	svc := service.NewAccountService(repo, logger)
	eps := endpoints.NewAccountEndpoint(svc)
	transport.NewHttpHandler(eps)
}
