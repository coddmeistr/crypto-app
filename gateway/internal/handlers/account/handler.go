package account_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/maxim12233/crypto-app-server/gateway"
	"github.com/maxim12233/crypto-app-server/gateway/internal/handlers"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
	account_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/account"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/jwt"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/matcher"

	"go.uber.org/zap"
)

const (
	accountGetAccountURL        = "/api/account/:id"
	accountCreateAccountURL     = "/api/account"
	accountDeleteAccountURL     = "/api/account/:id"
	accountGetAccountBalanceURL = "/api/account/:id/balance"
	accountBuyActivityURL       = "/api/account/:id/activity"
	accountSellActivityURL      = "/api/account/:id/activity"
	accountLoginURL             = "/api/account/login"
	accountFakeDepositURL       = "/api/account/:id/balance"
	accountGetActivitiesURL     = "/api/account/:id/activity"
)

type Handler struct {
	Logger         *zap.Logger
	AccountService account_service.IAccountService
}

func (h *Handler) Register(router *gin.Engine) {
	public := router.Group("")
	public.PUT(accountLoginURL, h.Login)
	public.POST(accountCreateAccountURL, h.CreateAccount)

	auth := router.Group("")
	auth.Use(
		jwt.AuthMiddleware([]uint{1}),
	)
	auth.GET(accountGetAccountURL, h.GetAccount)

	authWithMatcher := router.Group("")
	authWithMatcher.Use(
		jwt.AuthMiddleware([]uint{1}),
		matcher.MatchPathParamWithContextKey("id", "ID"),
	)
	authWithMatcher.DELETE(accountDeleteAccountURL, h.DeleteAccount)
	authWithMatcher.GET(accountGetAccountBalanceURL, h.GetAccountBalance)
	authWithMatcher.POST(accountBuyActivityURL, h.BuyActivity)
	authWithMatcher.DELETE(accountSellActivityURL, h.SellActivity)
	authWithMatcher.PUT(accountFakeDepositURL, h.FakeDeposit)
	authWithMatcher.GET(accountGetActivitiesURL, h.GetActivities)
}

func (h *Handler) GetActivities(c *gin.Context) {

	result, err := h.AccountService.GetActivities(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

func (h *Handler) FakeDeposit(c *gin.Context) {

	result, err := h.AccountService.FakeDeposit(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// GetAccount godoc
// @Summary Retrieve account info
// @Description Retrieves account basic info, based on given id
// @Tags  accounts
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User's account ID"
// @Success 200 {object} GetAccountResponse
// @Router /account/{id} [get]
func (h *Handler) GetAccount(c *gin.Context) {

	result, err := h.AccountService.GetAccount(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// CreateAccount godoc
// @Summary Creating new user account
// @Description Creating new unique user's account
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param Account_Info body string true "All required. Email must be valid" SchemaExample(login: NewUser\npassword: fsgdsfzg\nemail: euseew@yandex.ru)
// @Success 200 {object} string
// @Router /account [post]
func (h *Handler) CreateAccount(c *gin.Context) {

	result, err := h.AccountService.CreateAccount(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// DeleteAccount godoc
// @Summary Delete existing account
// @Description Delete user's account pernamently
// @Tags  accounts
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User's account ID"
// @Success 200 {object} string
// @Router /account/{id} [delete]
func (h *Handler) DeleteAccount(c *gin.Context) {

	result, err := h.AccountService.DeleteAccount(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// BuyActivity godoc
// @Summary Buy some crypto activity
// @Description Buys crypto activity for user for his balance
// @Tags  accounts
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User's account ID"
// @Param Buy_Info body string true "All required" SchemaExample(symbol: BTC\nprice: 1423)"
// @Success 200 {object} string
// @Router /account/{id}/activity [post]
func (h *Handler) BuyActivity(c *gin.Context) {

	result, err := h.AccountService.BuyActivity(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// SellActivity godoc
// @Summary Sell crypto activity
// @Description Sells existing user's crypto activity
// @Tags  accounts
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User's account ID"
// @Param Buy_Info body string true "All required" SchemaExample(symbol: BTC\nprice: 1423)"
// @Success 200 {object} string
// @Router /account/{id}/activity [delete]
func (h *Handler) SellActivity(c *gin.Context) {

	result, err := h.AccountService.SellActivity(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// GetBalance godoc
// @Summary Get user's balance
// @Description Get user's USD balance
// @Tags  accounts
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User's account ID"
// @Success 200 {object} GetAccountBalanceResponse
// @Router /account/{id}/balance [get]
func (h *Handler) GetAccountBalance(c *gin.Context) {

	result, err := h.AccountService.GetAccountBalance(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, result.Error)
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

// Login godoc
// @Summary Login user
// @Description If OK, payload contains authorization token(default exparation is ONE HOUR)
// @Description Every endpoint which requires authorization must contain token string in Authorization header
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param Login_Info body string true "Email required if login is empty and vice-versa" SchemaExample(login: SomeUser\npassword: fsgdsfzg\nemail: euseew@yandex.ru)
// @Success 200 {object} string
// @Router /account/login [put]
func (h *Handler) Login(c *gin.Context) {

	token, err := h.AccountService.Login(c)
	if err != nil {
		handlers.WriteJSONResponse(c, app.GetHTTPCodeFromError(err), nil, &models.Error{Code: app.ErrorCode(err), Message: err.Error()})
		return
	}

	handlers.WriteJSONResponse(c, http.StatusOK, token, nil)
}
