package crypto_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/maxim12233/crypto-app-server/gateway/internal/handlers"
	crypto_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/crypto"

	"go.uber.org/zap"
)

const (
	cryptoGetPricesURL  = "/api/crypto/prices"
	cryptoGetHistoryURL = "/api/crypto/history"
)

type Handler struct {
	Logger        *zap.Logger
	CryptoService crypto_service.ICryptoService
}

func (h *Handler) Register(router *gin.Engine) {
	router.GET(cryptoGetPricesURL, h.GetPrices)
	router.GET(cryptoGetHistoryURL, h.GetHistory)
}

func (h *Handler) GetPrices(c *gin.Context) {

	result, err := h.CryptoService.GetPrices(c)
	if err != nil {
		handlers.WriteJSONResponse(c, http.StatusBadGateway, nil, errors.New("Bad gateway"))
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, errors.New(result.Error.Message))
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}

func (h *Handler) GetHistory(c *gin.Context) {

	result, err := h.CryptoService.GetHistory(c)
	if err != nil {
		handlers.WriteJSONResponse(c, http.StatusBadGateway, nil, errors.New("Bad gateway"))
		return
	}
	if result.HaveError {
		handlers.WriteJSONResponse(c, result.HttpCode, nil, errors.New(result.Error.Message))
		return
	}

	handlers.WriteJSONResponse(c, result.HttpCode, result.Payload, nil)
}
