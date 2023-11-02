package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

type CryptoEndpoint struct {
	GetPrices          gin.HandlerFunc
	GetHistory         gin.HandlerFunc
	GetPriceDifference gin.HandlerFunc
	GetCandlesDataWs   gin.HandlerFunc
}

func NewCryptoEndpoint(s service.ICryptoService) CryptoEndpoint {
	eps := CryptoEndpoint{
		GetPrices:          MakeGetPricesEndpoint(s),
		GetHistory:         MakeGetHistoryEndpoint(s),
		GetPriceDifference: MakeGetTimePeriodPriceDifferenceEndpoint(s),
		GetCandlesDataWs:   MakeGetCandlesDataWebsocketEndpoint(s),
	}

	return eps
}
