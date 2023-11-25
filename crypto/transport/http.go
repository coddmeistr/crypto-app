package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/config"
	_ "github.com/maxim12233/crypto-app-server/crypto/docs"
	"github.com/maxim12233/crypto-app-server/crypto/endpoints"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(eps endpoints.CryptoEndpoint) *gin.Engine {
	router := gin.Default()
	c := config.GetConfig()
	if c.Environment.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("v1")
	{
		cryptoGroup := v1.Group("crypto")
		{
			cryptoGroup.GET("/prices", eps.GetPrices)
			cryptoGroup.GET("/history", eps.GetHistory)
			cryptoGroup.GET("/diff", eps.GetPriceDifference)
		}
	}

	return router
}
