package crypto_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/rest"
	"go.uber.org/zap"
)

const (
	getPricesResource  = "/prices"
	getHistoryResource = "/history"
)

type client struct {
	base rest.BaseClient
}

func NewService(baseURL string, logger *zap.Logger) ICryptoService {
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

type ICryptoService interface {
	GetPrices(ctx *gin.Context) (resp *models.Response, err error)
	GetHistory(ctx *gin.Context) (resp *models.Response, err error)
}

func (c *client) GetPrices(ctx *gin.Context) (*models.Response, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getPricesResource, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL. error: %v", err)
	}

	// Change destination URL and copy query from incoming request
	req := ctx.Request
	err = rest.ChangeRequestURLWithQuery(req, uri)
	if err != nil {
		return nil, err
	}

	apiresp, err := c.base.SendRequest(req)
	if err != nil {
		return nil, err
	}

	// Decode gained response to a basic response structure
	// It default struct for microservices in this application
	var resp models.Response
	if err = json.NewDecoder(apiresp.Response().Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *client) GetHistory(ctx *gin.Context) (*models.Response, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getHistoryResource, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL. error: %v", err)
	}

	// Change destination URL and copy query from incoming request
	req := ctx.Request
	err = rest.ChangeRequestURLWithQuery(req, uri)
	if err != nil {
		return nil, err
	}

	apiresp, err := c.base.SendRequest(req)
	if err != nil {
		return nil, err
	}

	// Decode gained response to a basic response structure
	// It default struct for microservices in this application
	var resp models.Response
	if err = json.NewDecoder(apiresp.Response().Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
