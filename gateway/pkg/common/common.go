package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/logger"
)

func ReturnAnauthorized(c *gin.Context) {
	logger := logger.GetLogger()
	logger.Error("Unauthorized access")
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Error: []ErrorDetail{
			{
				ErrorType:    "ErrorTypeUnauthorized",
				ErrorMessage: "You are not allowed to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}
