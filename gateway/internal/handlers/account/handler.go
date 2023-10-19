package account_handler

import (
	"github.com/gin-gonic/gin"

	account_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/account"

	"go.uber.org/zap"
)

const ()

type Handler struct {
	Logger         *zap.Logger
	AccountService account_service.IAccountService
}

func (h *Handler) Register(router *gin.Engine) {

}
