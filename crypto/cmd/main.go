package main

import (
	"flag"
	"fmt"

	"github.com/maxim12233/crypto-app-server/crypto/config"
	cryptocompare "github.com/maxim12233/crypto-app-server/crypto/crypto_compare_sdk"
	"github.com/maxim12233/crypto-app-server/crypto/endpoints"
	"github.com/maxim12233/crypto-app-server/crypto/repository"
	"github.com/maxim12233/crypto-app-server/crypto/service"
	"github.com/maxim12233/crypto-app-server/crypto/transport"
)

// @title Crypto Service API
// @version 1.0
// @description Swagger API for Golang Project Crypto Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email euseew.maxim2015@yandex.ru

// @license.name EUS

// @BasePath /v1/crypto
func main() {

	isDocker := flag.Bool("docker", false, "Defines if app runs with docker")
	flag.Parse()

	if *isDocker {
		config.Init("docker")
	} else {
		config.Init("local")
	}
	c := config.GetConfig()

	dbUrl := c.GetString("database.host")
	dbSession, err := repository.InitDB(dbUrl)
	if err != nil {
		panic(fmt.Errorf("Fatal error database connection: %s \n", err))
	}

	logger := config.InitializeLogger()

	market, err := cryptocompare.NewCryptoCompare("crypto-app", "8b0b296373b86403560d64ab9a1e29ae35ab088e87ef5745f554ef5d0a2673be")
	if err != nil {
		panic(fmt.Errorf("Fatal error market initialization: %s \n", err))
	}
	repo := repository.NewAccountRepository(dbSession, logger)
	svc := service.NewCryptoService(repo, logger, market)
	eps := endpoints.NewCryptoEndpoint(svc)
	transport.NewHttpHandler(eps)
}
