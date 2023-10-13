package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/account/service"
)

type AccountEndpoint struct {
	GetAccount    gin.HandlerFunc
	CreateAccount gin.HandlerFunc
	DeleteAccount gin.HandlerFunc
}

func NewAccountEndpoint(s service.IAccountService) AccountEndpoint {
	eps := AccountEndpoint{
		GetAccount:    MakeGetAccountEndpoint(s),
		DeleteAccount: MakeDeleteAccountEndpoint(s),
		CreateAccount: MakeCreateAccountEndpoint(s),
	}

	return eps
}
