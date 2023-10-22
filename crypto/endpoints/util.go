package endpoints

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	app "github.com/maxim12233/crypto-app-server/crypto"
	"github.com/maxim12233/crypto-app-server/crypto/models"
)

func notAllQueryError(notLeastedQueryName string) error {
	return app.WrapE(app.ErrNotAllRequiredQueries, fmt.Sprintf("Haven't got all required query params. First not leasted params with name: %s", notLeastedQueryName))
}

// Method checks all query params with "keys" names
// And if at least one doesn't exist it returns it name and "false"
// If all params exist it returns empty string and "true"
func hasRequiredQuery(q url.Values, keys ...string) (string, bool) {
	for _, v := range keys {
		if !q.Has(v) {
			return v, false
		}
	}
	return "", true
}

func writeJSONResponse(c *gin.Context, httpcode int, payload any, err error) {
	var e *models.Error
	if err != nil {
		e = &models.Error{
			Code:    app.ErrorCode(err),
			Message: err.Error(),
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
