package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/maxim12233/crypto-app-server/gateway/docs"
	"github.com/maxim12233/crypto-app-server/gateway/internal/config"
	account_handler "github.com/maxim12233/crypto-app-server/gateway/internal/handlers/account"
	crypto_handler "github.com/maxim12233/crypto-app-server/gateway/internal/handlers/crypto"
	account_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/account"
	crypto_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/crypto"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/cors"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/logger"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/metrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Crypto Service API
// @version 1.0
// @description Swagger API for Golang Project Crypto Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email euseew.maxim2015@yandex.ru

// @license.name EUS

// @securityDefinitions.basic  BearerAuth

// @host localhost:8282
// @BasePath /api
func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Couldn't load env variables from .env file")
	}

	_ = flag.Bool("docker", false, "Defines if app runs with docker")
	flag.Parse()

	logger := logger.Init()

	if err := config.Init("local"); err != nil {
		panic(err)
	}
	cfg := config.GetConfig()

	router := gin.Default()
	router.Use(cors.CORSMiddleware())

	// register swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	metricHandler := metrics.Metric{Logger: logger}
	metricHandler.Register(router)

	var cryptoBaseUrl = url.URL{
		Scheme: "http",
		Host:   cfg.CryptoService.Host,
		Path:   cfg.CryptoService.Url,
	}
	cryptoService := crypto_service.NewService(cryptoBaseUrl.String(), logger)
	cryptoHandler := crypto_handler.Handler{CryptoService: cryptoService, Logger: logger}
	cryptoHandler.Register(router)

	var accountBaseUrl = url.URL{
		Scheme: "http",
		Host:   cfg.AccountService.Host,
		Path:   cfg.AccountService.Url,
	}
	accountService := account_service.NewService(accountBaseUrl.String(), logger)
	accountHandler := account_handler.Handler{AccountService: accountService, Logger: logger}
	accountHandler.Register(router)

	start(router, logger)
}

func start(router *gin.Engine, logger *zap.Logger) {

	cfg := config.GetConfig()

	logger.Info(fmt.Sprintf("bind application to host: %s and port: %s", "localhost", cfg.Server.Port))

	/*listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", "localhost", cfg.GetString("server.port")))
	if err != nil {
		logger.Error(err.Error())
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Info("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
	*/
	logger.Info("application initialized and started")
	router.Run(cfg.Server.Port)
}
