package endpoints

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/service"
)

// GetPrice godoc
// @Summary Gets prices in choosen currencys for one crypto symbol
// @Produce json
// @Param symbol query string true "Crypto symbol: BTC"
// @Param symbolsTo query string true "All listed currencys which you want to see price for symbol"
// @Success 200 {object} crypto_compare_sdk.Prices
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
			writeJSONResponse(c, http.StatusBadGateway, nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, prices, nil)
	}
}

// GetHistory godoc
// @Summary Retrieves OHLCV info based on timebase, symbol, convert currency symbol
// @Produce json
// @Param timebase query string true "Timebase to get OHLCV data: days, hours, minutes"
// @Param symbol query string true "Main crypto symbol that you'll get OHLCV for: BTC"
// @Param symbolTo query string true "Price'd be converted to this currency symbol param"
// @Param limit query integer true "IMPORTANT: number of records you'll get"
// @Success 200 {object} crypto_compare_sdk.HistoricalData
// @Router /history [get]
func MakeGetHistoryEndpoint(s service.ICryptoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		if no, ok := hasRequiredQuery(q, "timebase", "symbol", "symbolTo", "limit"); !ok {
			writeJSONResponse(c, http.StatusBadRequest, nil, notAllQueryError(no))
			return
		}
		limit, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			writeJSONResponse(c, http.StatusBadRequest, nil, errInvalidParamType)
			return
		}

		data, err := s.GetHistory(q.Get("timebase"), q.Get("symbol"), q.Get("symbolTo"), limit)
		if err != nil {
			writeJSONResponse(c, http.StatusBadGateway, nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, data, nil)
	}
}

// GetHistory godoc
// @Summary Retrieves OHLCV info based on timebase, symbol, convert currency symbol
// @Produce json
// @Param timebase query string true "Timebase to get OHLCV data: days, hours, minutes"
// @Param symbol query string true "Main crypto symbol that you'll get OHLCV for: BTC"
// @Param symbolTo query string true "Price'd be converted to this currency symbol param"
// @Param limit query integer true "IMPORTANT: number of records you'll get"
// @Success 200 {object} crypto_compare_sdk.HistoricalData
// @Router /history [get]
func MakeGetTimePeriodPriceDifferenceEndpoint(s service.ICryptoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		if no, ok := hasRequiredQuery(q, "symbol", "symbolTo"); !ok {
			writeJSONResponse(c, http.StatusBadRequest, nil, notAllQueryError(no))
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
				writeJSONResponse(c, http.StatusBadRequest, nil, err)
				return
			}
		}

		data, err := s.GetPriceDifference(timebase, q.Get("symbol"), q.Get("symbolTo"), offset)
		if err != nil {
			writeJSONResponse(c, http.StatusBadGateway, nil, err)
			return
		}

		writeJSONResponse(c, http.StatusOK, data, nil)
	}
}
