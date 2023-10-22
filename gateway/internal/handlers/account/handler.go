package account_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/maxim12233/crypto-app-server/gateway"
	"github.com/maxim12233/crypto-app-server/gateway/internal/handlers"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
	account_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/account"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/jwt"

	"go.uber.org/zap"
)

const (
	accountGetAccountURL        = "/api/account/:id"
	accountCreateAccountURL     = "/api/account"
	accountDeleteAccountURL     = "/api/account/:id"
	accountGetAccountBalanceURL = "/api/account/:id/balance"
	accountBuyActivityURL       = "/api/account/:id/activity"
	accountLoginURL             = "/api/account/login"
	accountFakeDepositURL       = "/api/account/:id/balance"
	accountGetActivitiesURL     = "/api/account/:id/activity"
)

type Handler struct {
	Logger         *zap.Logger
	AccountService account_service.IAccountService
}

func (h *Handler) Register(router *gin.Engine) {
	router.GET(accountGetAccountURL, jwt.AuthMiddleware([]uint{1}), h.GetAccount)
	router.POST(accountCreateAccountURL, h.CreateAccount)
	router.DELETE(accountDeleteAccountURL, jwt.AuthMiddleware([]uint{1}), h.DeleteAccount)
	router.GET(accountGetAccountBalanceURL, jwt.AuthMiddleware([]uint{1}), h.GetAccountBalance)
	router.POST(accountBuyActivityURL, jwt.AuthMiddleware([]uint{1}), h.BuyActivity)
	router.PUT(accountLoginURL, h.Login)
	router.PUT(accountFakeDepositURL, jwt.AuthMiddleware([]uint{1}), h.FakeDeposit)
	router.GET(accountGetActivitiesURL, jwt.AuthMiddleware([]uint{1}), h.GetActivities)
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
