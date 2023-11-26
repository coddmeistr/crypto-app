package matcher

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/common"
	"github.com/maxim12233/crypto-app-server/gateway/pkg/logger"
)

func MatchPathParamWithContextKey(param string, key string) func(c *gin.Context) {
	return func(c *gin.Context) {

		p := c.Param(param)
		// Get key from context and convert it to string
		ivalue, ok := c.Get(key)
		if !ok {
			logger.GetLogger().Error(fmt.Sprintf("Key doesn't exist. Key name: %v", key))
			common.ReturnForbidden(c)
			return
		}
		fvalue, ok := ivalue.(float64)
		if !ok {
			logger.GetLogger().Error(fmt.Sprintf("Cannot convert key value to string. Failed interface{} to float64 convertion. Key name: %v", key))
			common.ReturnForbidden(c)
			return
		}
		svalue := fmt.Sprintf("%d", int(fvalue))

		// Compare matching
		if p != svalue {
			common.ReturnForbidden(c)
			return
		}

		c.Next()
	}
}
