package account_service

import (
	"net/http"
	"time"

	"github.com/maxim12233/crypto-app-server/gateway/pkg/rest"
	"go.uber.org/zap"
)

const ()

type client struct {
	base     rest.BaseClient
	Resource string
}

func NewService(baseURL string, logger *zap.Logger) IAccountService {
	c := client{
		base: rest.BaseClient{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger: logger,
		},
	}
	return &c
}

type IAccountService interface {
}
