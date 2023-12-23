package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/account/config"
	"github.com/maxim12233/crypto-app-server/account/endpoints"

	_ "github.com/maxim12233/crypto-app-server/account/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(eps endpoints.AccountEndpoint) *gin.Engine {
	router := gin.Default()

	c := config.GetConfig()
	if c.Environment.Mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)

	// swagger endp
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("v1")
	{
		accountGroup := v1.Group("account")
		{
			accountGroup.GET("/get", eps.GetAccount)
			accountGroup.GET("/:id", eps.GetAccount)
			accountGroup.DELETE("/:id", eps.DeleteAccount)
			accountGroup.GET("", eps.CreateAccount)
			accountGroup.GET("/login", eps.Login)
			accountGroup.GET("/getall", eps.GetAllAccounts)
			accountGroup.GET("/balance", eps.GetBalance)

			balance := accountGroup.Group("/:id/balance")
			balance.GET("", eps.GetBalance)
			balance.PUT("", eps.FakeDeposit)

			activity := accountGroup.Group("/:id/activity")
			activity.POST("", eps.BuyActivity)
			activity.DELETE("", eps.SellActivity)
			activity.GET("", eps.GetActivities)
		}
	}

	return router
}
