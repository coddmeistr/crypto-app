package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/internal/models"
)

func WriteJSONResponse(c *gin.Context, httpcode int, payload any, err *models.Error) {
	var e *models.Error
	if err != nil {
		e = &models.Error{
			Code:    err.Code,
			Message: err.Message,
		}
	} else {
		e = nil
	}
	c.JSON(httpcode, models.Response{
		HttpCode:  httpcode,
		HaveError: err != nil,
		Error:     e,
		Payload:   &payload,
	})
}
