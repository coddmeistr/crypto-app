package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	app "github.com/maxim12233/crypto-app-server/account"
	"github.com/maxim12233/crypto-app-server/account/models"
	"github.com/maxim12233/crypto-app-server/account/service"
)

// GetHistory godoc
// @Summary Retrieves OHLCV info based on timebase, symbol, convert currency symbol
// @Produce json
// @Param timebase query string true "Timebase to get OHLCV data: days, hours, minutes"
// @Param symbol query string true "Main crypto symbol that you'll get OHLCV for: BTC"
// @Param symbolTo query string true "Price'd be converted to this currency symbol param"
// @Param limit query integer true "IMPORTANT: number of records you'll get"
// @Success 200 {object} crypto_compare_sdk.HistoricalData
// @Router /history [get]
func MakeGetAccountEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		acc, err := s.GetAccount(uint(id))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, models.GetAccountResponse{
			ID:    acc.ID,
			Login: acc.Login,
			Email: acc.Email,
		}, nil)
	}
}

func MakeCreateAccountEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.CreateAccountRequest
		if err := c.BindJSON(&body); err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse json body"))
			return
		}

		// Validation
		v := validator.New()
		if err := v.Struct(body); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field(), err.Tag())
			}
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrValidation), nil, app.WrapE(app.ErrValidation, "Body contains invalid data"))
			return
		}

		err := s.CreateAccount(body)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusCreated, "New account successfully created.", nil)
	}
}

func MakeDeleteAccountEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		err = s.DeleteAccount(uint(id))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, "Account was successfully deleted forever.", nil)
	}
}

func MakeLoginEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.LoginRequest
		if err := c.BindJSON(&body); err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse json body"))
			return
		}

		// Validation
		v := validator.New()
		if err := v.Struct(body); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field(), err.Tag())
			}
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrValidation), nil, app.WrapE(app.ErrValidation, "Body contains invalid data"))
			return
		}

		id, err := s.Login(body)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, models.LoginResponse{UserID: id}, nil)
	}
}

func MakeGetAccountBalanceEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		bal, err := s.GetBalance(uint(id))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, models.GetAccountBalanceResponse{
			AccountID: bal.AccountID,
			USD:       *bal.USD,
		}, nil)
	}
}

func MakeBuyActivityEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		var body models.BuyActivityRequest
		if err := c.BindJSON(&body); err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse json body"))
			return
		}

		// Validation
		v := validator.New()
		if err = v.Struct(body); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Field(), err.Tag())
			}
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrValidation), nil, app.WrapE(app.ErrValidation, "Body contains invalid data"))
			return
		}

		err = s.BuyActivity(uint(id), body.Symbol, body.Price)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, "Activity successfully created.", nil)
	}
}
