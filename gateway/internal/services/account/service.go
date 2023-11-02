package account_service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	app "github.com/maxim12233/crypto-app-server/gateway"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/jwt"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/rest"
	"go.uber.org/zap"
)

const (
	getAccountResource        = "/:id"
	createAccountResource     = ""
	deleteAccountResource     = "/:id"
	getAccountBalanceResource = "/:id/balance"
	buyActivityResource       = "/:id/activity"
	sellActivityResource      = "/:id/activity"
	loginResource             = "/login"
	fakeDepositResource       = "/:id/balance"
	getActivitiesResource     = "/:id/activity"
)

type client struct {
	base   rest.BaseClient
	logger *zap.Logger
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
		logger: logger,
	}
	return &c
}

type IAccountService interface {
	GetAccount(ctx *gin.Context) (*models.Response, error)
	DeleteAccount(ctx *gin.Context) (*models.Response, error)
	CreateAccount(ctx *gin.Context) (*models.Response, error)
	GetAccountBalance(ctx *gin.Context) (*models.Response, error)
	BuyActivity(ctx *gin.Context) (*models.Response, error)
	SellActivity(ctx *gin.Context) (*models.Response, error)
	Login(ctx *gin.Context) (*string, error)
	FakeDeposit(ctx *gin.Context) (*models.Response, error)
	GetActivities(ctx *gin.Context) (*models.Response, error)
}

func (c *client) GetActivities(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getActivitiesResource, nil, pathParams)
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

func (c *client) FakeDeposit(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(fakeDepositResource, nil, pathParams)
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

func (c *client) GetAccount(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getAccountResource, nil, pathParams)
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

func (c *client) DeleteAccount(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(deleteAccountResource, nil, pathParams)
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

func (c *client) CreateAccount(ctx *gin.Context) (*models.Response, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(createAccountResource, nil, nil)
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

func (c *client) GetAccountBalance(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(getAccountBalanceResource, nil, pathParams)
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

func (c *client) BuyActivity(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(buyActivityResource, nil, pathParams)
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

func (c *client) SellActivity(ctx *gin.Context) (*models.Response, error) {

	pathParams := rest.PathOptions{
		"id": ctx.Param("id"),
	}

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(sellActivityResource, nil, pathParams)
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

func (c *client) Login(ctx *gin.Context) (*string, error) {

	// Build resource URL: it's basically baseURL + resourceURL
	uri, err := c.base.BuildURL(loginResource, nil, nil)
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
	req.Method = http.MethodGet

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

	// Various error handlings
	if resp.HaveError {
		err, ok := app.CodeToError(resp.Error.Code)
		if !ok {
			c.logger.Error("Error gained from account service is undefined. Probably disturbed error codes standart")
			return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
		}
		switch err {
		case app.ErrNotFound:
			c.logger.Error("Not found login or email", zap.Error(err))
			return nil, app.WrapE(err, "Requested account doesn't exist")
		case app.ErrIncorrectLoginOrPassword:
			c.logger.Error("Incorrect login or password", zap.Error(err))
			return nil, app.WrapE(app.ErrIncorrectLoginOrPassword, "Invalid login or password(actually login is valid, because if it doesn't exist it throws Not Found)")
		}

		return nil, err
	}

	// According to account microservice documantation
	// Payload should store int value in one struct with name id
	payload := *resp.Payload
	payloadMap, ok := payload.(map[string]interface{})
	if !ok {
		c.logger.Error("Invalid payload data from login endpoint: with payloadMap", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}
	id, ok := payloadMap["id"].(float64)
	if !ok {
		c.logger.Error("Invalid payload data from login endpoint: with id", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with some of the service")
	}

	// Generate new jwt
	jwthelper := jwt.NewHelper(c.logger)
	token, err := jwthelper.GenerateJWT(uint(id), []uint{1})
	if err != nil {
		c.logger.Error("Error while generating jwt token", zap.Error(err))
		return nil, app.WrapE(app.ErrInternal, "Error with authorization")
	}

	return &token, nil
}
