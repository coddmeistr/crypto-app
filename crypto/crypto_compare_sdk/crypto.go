package cryptocompare

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	host                 = "https://min-api.cryptocompare.com"
	pricesURL            = "/data/price"
	historicalDaysURL    = "/data/v2/histoday"
	historicalHoursURL   = "/data/v2/histohour"
	historicalMinutesURL = "/data/v2/histominute"
)

var (
	errEmptyApiKey    = errors.New("Cryptocompare API key cannot be empty")
	errInvalidHostURL = errors.New("Cryptocompare sdk contains invalid host url")
	errCodeNotOK      = errors.New("Something went wrong on the foreign API side")
)

type ICryptoCompare interface {
	GetLatestPrice(symbol string, symbolTo []string) (*Prices, error)
	GetHistoricalDailyOHLCV(fsym string, tsym string, limit int) (*HistoricalData, error)
	GetHistoricalHourlyOHLCV(fsym string, tsym string, limit int) (*HistoricalData, error)
	GetHistoricalMinutlyOHLCV(fsym string, tsym string, limit int) (*HistoricalData, error)
}

type CryptoCompare struct {
	appName string
	apiKey  string
	client  *http.Client
}

func NewCryptoCompare(appName string, apiKey string) (ICryptoCompare, error) {
	if apiKey == "" {
		return nil, errEmptyApiKey
	}
	_, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, errInvalidHostURL
	}
	return &CryptoCompare{
		appName: appName,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}, nil
}

func (c *CryptoCompare) buildBaseURL(resource string) url.URL {
	url, _ := url.ParseRequestURI(host)
	url.Path = path.Join(url.Path, resource)

	q := url.Query()
	q.Set("api_key", c.apiKey)
	q.Set("extraParams", c.appName)
	url.RawQuery = q.Encode()

	return *url
}

func (c *CryptoCompare) GetLatestPrice(symbol string, symbolsTo []string) (*Prices, error) {

	url := c.buildBaseURL(pricesURL)

	symbols := strings.Join(symbolsTo, ",")

	q := url.Query()
	q.Set("fsym", symbol)
	q.Set("tsyms", symbols)
	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errCodeNotOK
	}

	var rm map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&rm); err != nil {
		return nil, err
	}

	return &Prices{
		Prices: rm,
	}, nil
}

func (c *CryptoCompare) GetHistoricalDailyOHLCV(fsym string, tsym string, limit int) (*HistoricalData, error) {

	url := c.buildBaseURL(historicalDaysURL)

	q := url.Query()
	q.Set("fsym", fsym)
	q.Set("tsym", tsym)
	q.Set("limit", fmt.Sprint(limit))
	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errCodeNotOK
	}

	var r HistoricalResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r.HistoricalData, nil
}

func (c *CryptoCompare) GetHistoricalHourlyOHLCV(fsym string, tsym string, limit int) (*HistoricalData, error) {

	url := c.buildBaseURL(historicalHoursURL)

	q := url.Query()
	q.Set("fsym", fsym)
	q.Set("tsym", tsym)
	q.Set("limit", fmt.Sprint(limit))
	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errCodeNotOK
	}

	var r HistoricalResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r.HistoricalData, nil
}

func (c *CryptoCompare) GetHistoricalMinutlyOHLCV(fsym string, tsym string, limit int) (*HistoricalData, error) {

	url := c.buildBaseURL(historicalMinutesURL)

	q := url.Query()
	q.Set("fsym", fsym)
	q.Set("tsym", tsym)
	q.Set("limit", fmt.Sprint(limit))
	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errCodeNotOK
	}

	var r HistoricalResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r.HistoricalData, nil
}
