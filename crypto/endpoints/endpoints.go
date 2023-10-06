package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

// GetCurrency godoc
// @Summary Retrieves account based on given ID
// @Produce json
// @Param id path integer true "Currency ID"
// @Success 200 {object} models.Currency
// @Router /{id} [get]
func MakeGetAccountEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		currency, err := s.GetAccountInfoById(uint(id))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"currency": currency,
		})
	}
}
