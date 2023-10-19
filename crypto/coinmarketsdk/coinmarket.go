package coinmarket

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
)

const (
	host      = "https://pro-api.coinmarketcap.com"
	quotesURL = "/v2/cryptocurrency/quotes/latest"
)

type ICoinMarket interface {
	GetLatestQuotes(slug string) (*USD, error)
}

type CoinMarket struct {
	apikey string
	client *http.Client
}

func NewCoinMarket(key string) ICoinMarket {
	return &CoinMarket{
		apikey: key,
		client: http.DefaultClient,
	}
}

func (c *CoinMarket) GetLatestQuotes(slug string) (*USD, error) {
	url, err := url.ParseRequestURI(host)
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(url.Path, quotesURL)

	q := url.Query()
	q.Set("slug", slug)
	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", c.apikey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Something is not ok")
	}

	var r QuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return nil, nil
}
