package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	app "github.com/maxim12233/crypto-app-server/account"
	"github.com/maxim12233/crypto-app-server/account/config"
	"github.com/maxim12233/crypto-app-server/account/models"
	"go.uber.org/zap"
)

// Helper function to make requests to foreign microservice
// This function fetches current price for given symbol in given currency
// Logger instance is required due to many error logic
func fetchSymbolPriceFromCryptoMicroservice(symbol string, currency string, logger *zap.Logger) (float64, error) {

	// Take the current price for symbol via asking crypto microservice
	// Using config values to get host and path to foreign microservice
	cfg := config.GetConfig()
	var currentSymbolPrice float64
	host := cfg.Dependencies.CryptoService.Host
	resPath := cfg.Dependencies.CryptoService.Endpoints.GetCurrentPrices
	uri := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   resPath,
	}
	q := uri.Query()
	q.Set("symbol", symbol)
	q.Set("symbolsTo", currency)
	uri.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		logger.Error("Couldn't create new request to request crypto service", zap.Error(err))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Error while doing request to foreign api", zap.Error(err))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error("Foreign api, got not 200 code response")
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	var body models.Response // Basic response model for all of this app's microservices and gateway
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		logger.Error("Error while decoding json from crypto service response", zap.Error(err))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	// Get the payload response struct
	var payload interface{}
	if body.Payload == nil {
		logger.Error("Error, got an empty nil payload from crypto service", zap.Error(err))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	payload = *body.Payload
	pricesMap, ok := payload.(map[string]interface{})
	if !ok {
		logger.Error("Cannot cast payload to proper response type", zap.Error(err))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}
	prices, ok := pricesMap["Prices"].(map[string]interface{})
	if !ok {
		logger.Error("Cannot cast payload to proper response type", zap.Error(err))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	// Check if needed value exists
	if _, ok := prices[currency]; !ok {
		logger.Error(fmt.Sprintf("Final map dont contain %s key", currency))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	// Get this value and try to covert to wanted type
	currentSymbolPrice, ok = prices[currency].(float64)
	if !ok {
		logger.Error(fmt.Sprintf("Couldn't convert price to float64"))
		return 0, app.WrapE(app.ErrInternal, "Foreign server issue or internal problem")
	}

	return currentSymbolPrice, nil
}
