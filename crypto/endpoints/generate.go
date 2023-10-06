package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

type AccountEndpoint struct {
	GetAccount gin.HandlerFunc
}

func NewAccountEndpoint(s service.IAccountService) AccountEndpoint {
	eps := AccountEndpoint{
		GetAccount: MakeGetAccountEndpoint(s),
	}

	return eps
}
