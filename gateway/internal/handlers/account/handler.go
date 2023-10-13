package account_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	account_service "github.com/maxim12233/crypto-app-server/gateway/internal/services/account"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/jwt"

	"go.uber.org/zap"
)

const (
	accountURL = "/api/account/:id"
)

type Handler struct {
	Logger         *zap.Logger
	AccountService account_service.IAccountService
}

func (h *Handler) Register(router *gin.Engine) {
	router.GET(accountURL, jwt.AuthMiddleware([]uint{1, 2}), h.GetAccount)
}

func (h *Handler) GetAccount(c *gin.Context) {

	var a account_service.Account

	a, err := h.AccountService.GetAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  1,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account": a,
		"code":    0,
	})
}
