package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	bitfinex "github.com/maxim12233/crypto-app-server/crypto/bitfinex_sdk"
	"github.com/maxim12233/crypto-app-server/crypto/config"
	cryptocompare "github.com/maxim12233/crypto-app-server/crypto/crypto_compare_sdk"
	"github.com/maxim12233/crypto-app-server/crypto/endpoints"
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
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Couldn't load env variables from .env file")
	}

	_ = flag.Bool("docker", false, "Defines if app runs with docker")
	flag.Parse()

	if err := config.Init("local"); err != nil {
		panic(err)
	}
	c := config.GetConfig()

	logger := config.InitializeLogger()

	market, err := cryptocompare.NewCryptoCompare(c.GetString("crypto_compare.app_name"), os.Getenv("CRYPTO_COMPARE_KEY"))
	if err != nil {
		panic(fmt.Errorf("Fatal error market initialization: %s \n", err))
	}

	wsmarket, err := bitfinex.NewBitfinex()
	if err != nil {
		panic(fmt.Errorf("Fatal error wsmarket initialization: %s \n", err))
	}

	svc := service.NewCryptoService(logger, market, wsmarket)
	eps := endpoints.NewCryptoEndpoint(svc)
	transport.NewHttpHandler(eps)
}
