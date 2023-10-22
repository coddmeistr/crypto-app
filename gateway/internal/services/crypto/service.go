package crypto_service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	app "github.com/maxim12233/crypto-app-server/gateway"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/rest"
	"go.uber.org/zap"
)

const (
	getPricesResource     = "/prices"
	getHistoryResource    = "/history"
	getDifferenceResource = "/diff"
)

type client struct {
	base   rest.BaseClient
	logger *zap.Logger
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
		logger: logger,
	}
	return &c
}

type ICryptoService interface {
	GetPrices(ctx *gin.Context) (resp *models.Response, err error)
	GetHistory(ctx *gin.Context) (resp *models.Response, err error)
	GetDifference(ctx *gin.Context) (resp *models.Response, err error)
}

func (c *client) GetPrices(ctx *gin.Context) (*models.Response, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getPricesResource, nil, nil)
	if err != nil {
		c.logger.Error("Error while building url", zap.Error(err))
		return nil, app.ErrInternal
	}

	// Change destination URL and copy query from incoming request
	req := ctx.Request
	err = rest.ChangeRequestURLWithQuery(req, uri)
	if err != nil {
		c.logger.Error("Error while copying query", zap.Error(err))
		return nil, app.ErrInternal
	}

	apiresp, err := c.base.SendRequest(req)
	if err != nil {
		c.logger.Error("Error sending request", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	// Decode gained response to a basic response structure
	// It default struct for microservices in this application
	var resp models.Response
	if err = json.NewDecoder(apiresp.Response().Body).Decode(&resp); err != nil {
		c.logger.Error("Error decoding json from response body", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	return &resp, nil
}

func (c *client) GetHistory(ctx *gin.Context) (*models.Response, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getHistoryResource, nil, nil)
	if err != nil {
		c.logger.Error("Error while building url", zap.Error(err))
		return nil, app.ErrInternal
	}

	// Change destination URL and copy query from incoming request
	req := ctx.Request
	err = rest.ChangeRequestURLWithQuery(req, uri)
	if err != nil {
		c.logger.Error("Error while copying query", zap.Error(err))
		return nil, app.ErrInternal
	}

	apiresp, err := c.base.SendRequest(req)
	if err != nil {
		c.logger.Error("Error sending request", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	// Decode gained response to a basic response structure
	// It default struct for microservices in this application
	var resp models.Response
	if err = json.NewDecoder(apiresp.Response().Body).Decode(&resp); err != nil {
		c.logger.Error("Error decoding json from response body", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	return &resp, nil
}

func (c *client) GetDifference(ctx *gin.Context) (*models.Response, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getDifferenceResource, nil, nil)
	if err != nil {
		c.logger.Error("Error while building url", zap.Error(err))
		return nil, app.ErrInternal
	}

	// Change destination URL and copy query from incoming request
	req := ctx.Request
	err = rest.ChangeRequestURLWithQuery(req, uri)
	if err != nil {
		c.logger.Error("Error while copying query", zap.Error(err))
		return nil, app.ErrInternal
	}

	apiresp, err := c.base.SendRequest(req)
	if err != nil {
		c.logger.Error("Error sending request", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	// Decode gained response to a basic response structure
	// It default struct for microservices in this application
	var resp models.Response
	if err = json.NewDecoder(apiresp.Response().Body).Decode(&resp); err != nil {
		c.logger.Error("Error decoding json from response body", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	return &resp, nil
}
