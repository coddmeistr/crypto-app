package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
)

func WriteJSONResponse(c *gin.Context, httpcode int, payload any, err error) {
	var e *models.Error
	if err != nil {
		e = &models.Error{
			Message: err.Error(),
		}
	} else {
		e = nil
	}
	c.JSON(httpcode, models.Response{
		HttpCode:  httpcode,
		HaveError: err != nil,
		Error:     e,
		Payload:   payload,
	})
}
