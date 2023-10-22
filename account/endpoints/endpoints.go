package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	app "github.com/maxim12233/crypto-app-server/account"
	"github.com/maxim12233/crypto-app-server/account/service"
)

// GetAccount godoc
// @Summary Retrieve account info
// @Description Retrieves account basic info, based on given id
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param id path int true "User's account ID"
// @Success 200 {object} models.GetAccountResponse
// @Router /{id} [get]
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

		writeJSONResponse(c, http.StatusOK, GetAccountResponse{
			ID:    acc.ID,
			Login: acc.Login,
			Email: acc.Email,
		}, nil)
	}
}

// CreateAccount godoc
// @Summary Creating new user account
// @Description Creating new unique user's account
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param login body string true "User's unique login"
// @Param password body string true "User's password"
// @Param email body string true "Valid unique email"
// @Success 200 {object} string
// @Router / [post]
func MakeCreateAccountEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body CreateAccountRequest
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

		err := s.CreateAccount(body.Login, body.Password, body.Email)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusCreated, "New account successfully created.", nil)
	}
}

// DeleteAccount godoc
// @Summary Delete existing account
// @Description Delete user's account pernamently
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param id path int true "User's account ID"
// @Success 200 {object} string
// @Router /{id} [delete]
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

// Login godoc
// @Summary Login user
// @Description Check if user can login. Response contains user's id
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param login body string false "User's login, required if email param is empty"
// @Param password body string true "User's password"
// @Param email body string false "User's email, required if login field is empty"
// @Success 200 {object} models.LoginResponse
// @Router /login [get]
func MakeLoginEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body LoginRequest
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

		id, err := s.Login(body.Login, body.Password, body.Email)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, LoginResponse{UserID: id}, nil)
	}
}

// GetBalance godoc
// @Summary Get user's balance
// @Description Get user's USD balance
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param id path int true "User's account ID"
// @Success 200 {object} models.GetAccountBalanceResponse
// @Router /{id}/balance [get]
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

		writeJSONResponse(c, http.StatusOK, GetAccountBalanceResponse{
			AccountID: bal.AccountID,
			USD:       *bal.USD,
		}, nil)
	}
}

func MakeFakeDepositEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		no, ok := hasRequiredQuery(c.Request.URL.Query(), "deposit")
		if !ok {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, notAllQueryError(no))
			return
		}
		deposit, err := strconv.ParseFloat(c.Query("deposit"), 10)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "deposit param has invalid type"))
			return
		}

		err = s.FakeDeposit(uint(id), deposit)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, "Deposit succesessful", nil)
	}
}

func MakeGetUserActivitiesEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		symbols := c.Query("symbols")

		records, err := s.GetActivities(uint(id), symbols)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}
		var activities []ActivityResponse
		for _, v := range records {
			activities = append(activities, ActivityResponse{Symbol: v.Symbol, Amount: v.Amount})
		}
		resp := GetActivitiesResponse{
			AccountID:  uint(id),
			Activities: activities,
		}

		writeJSONResponse(c, http.StatusOK, resp, nil)
	}
}

// BuyActivity godoc
// @Summary Buy some crypto activity
// @Description Buys crypto activity for user for his balance
// @Tags  accounts
// @Accept  json
// @Produce json
// @Param symbol body string true "Price string tag. Example: BTC"
// @Param price body string true "How much money user want to spend on this cryptocurrency in USD"
// @Success 200 {object} string
// @Router /{id}/activity [post]
func MakeBuyActivityEndpoint(s service.IAccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'id' path param"))
			return
		}

		var body BuyActivityRequest
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
