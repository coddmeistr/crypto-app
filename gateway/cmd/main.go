package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/internal/config"
	account_handler "github.com/maxim12233/crypto-app-server/gateway/internal/handlers/account"
	account_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/account"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/logger"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/metrics"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger := logger.Init()

	config.Init("local")
	cfg := config.GetConfig()

	router := gin.Default()

	metricHandler := metrics.Metric{Logger: logger}
	metricHandler.Register(router)

	accountService := account_service.NewService(cfg.GetString("account_service.url"), "/account/:id", logger)
	accountHandler := account_handler.Handler{AccountService: accountService, Logger: logger}
	accountHandler.Register(router)

	start(router, logger, cfg)
}

func start(router *gin.Engine, logger *zap.Logger, cfg *viper.Viper) {

	logger.Info(fmt.Sprintf("bind application to host: %s and port: %s", "localhost", cfg.GetString("server.port")))

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
	router.Run(cfg.GetString("server.port"))
}
