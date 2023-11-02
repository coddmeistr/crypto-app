package crypto_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/maxim12233/crypto-app-server/gateway"
	"github.com/maxim12233/crypto-app-server/gateway/internal/config"
	"github.com/maxim12233/crypto-app-server/gateway/internal/handlers"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
	crypto_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/crypto"

	"go.uber.org/zap"
)

const (
	cryptoGetPricesURL     = "/api/crypto/prices"
	cryptoGetHistoryURL    = "/api/crypto/history"
	cryptoGetDifferenceURL = "/api/crypto/diff"

	cryptoWebsocketConnectionURL = "/ws/connect"
)

type Handler struct {
	Logger        *zap.Logger
	CryptoService crypto_service.ICryptoService
}

func (h *Handler) Register(router *gin.Engine) {
	router.GET(cryptoGetPricesURL, h.GetPrices)
	router.GET(cryptoGetHistoryURL, h.GetHistory)
	router.GET(cryptoGetDifferenceURL, h.GetDifference)
	router.GET(cryptoWebsocketConnectionURL, h.EstablishWebsocketConnection)
}

// GetPrices godoc
// @Summary Get latest prices
// @Description Retrieves latest prices in different currencies for given crypto symbol
// @Description Every field in Prices response object is your given "symbolsTo" value.
// @Tags  crypto
// @Accept  json
// @Produce json
// @Param symbol query string true "Crypto currency symbol. Example: BTC"
// @Param symbolsTo query string true "One or many default currencies to convert crypto symbol to. Example 1: USD | Example 2: USD,JPY"
// @Success 200 {object} Prices
// @Router /crypto/prices [get]
func (h *Handler) GetPrices(c *gin.Context) {

	result, err := h.CryptoService.GetPrices(c)
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

// GetHistory godoc
// @Summary Get OHLCV history info
// @Description Getting Open High Low Close Volume info about given symbol
// @Description Uses different timebases depends on timebase query param
// @Tags  crypto
// @Accept  json
// @Produce json
// @Param timebase query string true "What time you want to track. All variants: days, hours, minutes"
// @Param symbol query string true "Crypto currency symbol. Example: BTC"
// @Param symbolTo query string true "One default currency to convert crypto symbol to. Example: USD"
// @Param limit query int true "How many records you want to get. For example: timebase=days limit=5 means that you get 5 days history from current date"
// @Success 200 {object} HistoricalData
// @Router /crypto/history [get]
func (h *Handler) GetHistory(c *gin.Context) {

	result, err := h.CryptoService.GetHistory(c)
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

// GetTimePeriodPriceDifference godoc
// @Summary Price difference
// @Description Getting price difference in USD and % between current date and some historical date
// @Description Use query params to configure it right
// @Tags  crypto
// @Accept  json
// @Produce json
// @Param timebase query string false "Default: days. What time you want to track. All variants: days, hours, minutes"
// @Param symbol query string true "Crypto currency symbol. Example: BTC"
// @Param symbolTo query string true "One default currency to convert crypto symbol to. Example: USD"
// @Param offset query int false "Default: 1. Offset from current date. For example if timebase=days and offset=3 you get price difference between current day and day that was 3 days ago"
// @Success 200 {object} PriceDifference
// @Router /crypto/diff [get]
func (h *Handler) GetDifference(c *gin.Context) {

	result, err := h.CryptoService.GetDifference(c)
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

func (h *Handler) EstablishWebsocketConnection(c *gin.Context) {
	wsUrl := config.GetConfig().GetString("crypto_service.ws")

	c.Header("Location", wsUrl)

	handlers.WriteJSONResponse(c, http.StatusTemporaryRedirect, "Resource temporary moved to other URL", nil)
}
