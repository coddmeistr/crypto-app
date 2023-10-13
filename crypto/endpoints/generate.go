package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

type CryptoEndpoint struct {
	GetQuote gin.HandlerFunc
}

func NewCryptoEndpoint(s service.ICryptoService) CryptoEndpoint {
	eps := CryptoEndpoint{
		GetQuote: MakeGetQuoteEndpoint(s),
	}

	return eps
}
