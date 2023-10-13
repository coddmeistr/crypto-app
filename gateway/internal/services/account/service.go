package account_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/rest"
	"go.uber.org/zap"
)

type client struct {
	base     rest.BaseClient
	Resource string
}

func NewService(baseURL string, resource string, logger *zap.Logger) IAccountService {
	c := client{
		Resource: resource,
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
	GetAccount(ctx *gin.Context) (Account, error)
}

func (c *client) GetAccount(ctx *gin.Context) (a Account, err error) {
	c.base.Logger.Debug("Generating query and path parameters")
	filters := []rest.FilterOptions{}
	pathparams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	c.base.Logger.Debug("build url with resource")
	uri, err := c.base.BuildURL(c.Resource, filters, pathparams)
	if err != nil {
		return a, fmt.Errorf("failed to build URL. error: %v", err)
	}

	c.base.Logger.Debug("create new request")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return a, fmt.Errorf("failed to create new request due to error: %w", err)
	}

	c.base.Logger.Debug("send request")
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(reqCtx)
	response, err := c.base.SendRequest(req)
	if err != nil {
		return a, fmt.Errorf("failed to send request due to error: %w", err)
	}

	if response.IsOk {
		if err = json.NewDecoder(response.Body()).Decode(&a); err != nil {
			return a, fmt.Errorf("failed to decode body due to error %w", err)
		}
		return a, nil
	}

	return a, errors.New("Gateway error")
}
