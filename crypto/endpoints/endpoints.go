package endpoints

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	app "github.com/maxim12233/crypto-app-server/crypto"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

// GetPrices godoc
// @Summary Get latest prices
// @Description Retrieves latest prices in different currencies for given crypto symbol
// @Description Every field in Prices response object is your given "symbolsTo" value.
// @Tags  crypto
// @Accept  json
// @Produce json
// @Param symbol query string true "Crypto currency symbol. Example: BTC"
// @Param symbolsTo query string true "One or many default currencies to convert crypto symbol to. Example 1: USD | Example 2: USD,JPY"
// @Success 200 {object} cryptocompare.Prices
// @Router /prices [get]
func MakeGetPricesEndpoint(s service.ICryptoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		if no, ok := hasRequiredQuery(q, "symbol", "symbolsTo"); !ok {
			writeJSONResponse(c, http.StatusBadRequest, nil, notAllQueryError(no))
			return
		}

		prices, err := s.GetPrice(q.Get("symbol"), strings.Split(q.Get("symbolsTo"), ","))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, prices, nil)
	}
}

// GetHistory godoc
// @Summary Get OHLCV history info
// @Description Getting Open High Low Close Volume info about given symbol
// @Description Uses different timebases depends on timebase query param
// @Tags  crypto
// @Accept  json
// @Produce json
// @Param timebase query string true "What time you want to track. All variants: days, hours, minutes"
// @Param symbol query string true "Crypto currency symbol. Example: BTC"
// @Param symbolTo query string true "One default currency to convert crypto symbol to. Example: USD"
// @Param limit query int true "How many records you want to get. For example: timebase=days limit=5 means that you get 5 days history from current date"
// @Success 200 {object} cryptocompare.HistoricalData
// @Router /history [get]
func MakeGetHistoryEndpoint(s service.ICryptoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		if no, ok := hasRequiredQuery(q, "timebase", "symbol", "symbolTo", "limit"); !ok {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, notAllQueryError(no))
			return
		}
		limit, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'limit' query param"))
			return
		}

		data, err := s.GetHistory(q.Get("timebase"), q.Get("symbol"), q.Get("symbolTo"), limit)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, data, nil)
	}
}

// GetTimePeriodPriceDifference godoc
// @Summary Price difference
// @Description Getting price difference in USD and % between current date and some historical date
// @Description Use query params to configure it right
// @Tags  crypto
// @Accept  json
// @Produce json
// @Param timebase query string false "Default: days. What time you want to track. All variants: days, hours, minutes"
// @Param symbol query string true "Crypto currency symbol. Example: BTC"
// @Param symbolTo query string true "One default currency to convert crypto symbol to. Example: USD"
// @Param offset query int false "Default: 1. Offset from current date. For example if timebase=days and offset=3 you get price difference between current day and day that was 3 days ago"
// @Success 200 {object} models.PriceDifference
// @Router /diff [get]
func MakeGetTimePeriodPriceDifferenceEndpoint(s service.ICryptoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		if no, ok := hasRequiredQuery(q, "symbol", "symbolTo"); !ok {
			writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, notAllQueryError(no))
			return
		}

		var timebase string
		if !q.Has("timebase") {
			timebase = "days"
		}
		var offset int
		var err error
		if !q.Has("offset") {
			offset = 1
		} else {
			offset, err = strconv.Atoi(q.Get("offset"))
			if err != nil {
				writeJSONResponse(c, app.GetHTTPCodeFromError(app.ErrBadRequest), nil, app.WrapE(app.ErrBadRequest, "Couldn't parse 'offset' query param"))
				return
			}
		}

		data, err := s.GetPriceDifference(timebase, q.Get("symbol"), q.Get("symbolTo"), offset)
		if err != nil {
			writeJSONResponse(c, app.GetHTTPCodeFromError(err), nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, data, nil)
	}
}
