package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

// GetCurrency godoc
// @Summary Retrieves account based on given ID
// @Produce json
// @Param id path integer true "Currency ID"
// @Success 200 {object} models.Currency
// @Router /{id} [get]
func MakeGetQuoteEndpoint(s service.ICryptoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Params.ByName("slug")
		quote, err := s.GetQuote(slug)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"quote": quote,
		})
	}
}
