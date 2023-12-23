package rest

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"go.uber.org/zap"
)

type BaseClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Logger     *zap.Logger
}

func (c *BaseClient) SendRequest(req *http.Request) (*APIResponse, error) {
	if c.HTTPClient == nil {
		return nil, errors.New("no http client")
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request. error: %w", err)
	}

	apiResponse := APIResponse{
		response: response,
	}

	return &apiResponse, nil
}

func (c *BaseClient) BuildURL(resource string, filters []FilterOptions, pathparams PathOptions) (string, error) {
	var resultURL string
	parsedURL, err := url.ParseRequestURI(c.BaseURL)
	if err != nil {
		return resultURL, fmt.Errorf("failed to parse base URL. error: %w", err)
	}

	// Parsing and putting path params to the destination url
	if pathparams != nil {

		urlwords := strings.Split(resource, "/")
		for i, v := range urlwords {
			if len(v) != 0 && v[0] == ':' {
				if param, ok := pathparams[v[1:]]; ok {
					urlwords[i] = param
				} else {
					return resultURL, fmt.Errorf("failed to parse path parameters. error: %w", err)
				}
			}
		}
		resource = strings.Join(urlwords, "/")
	}

	// Getting final url with path params and host
	parsedURL.Path = path.Join(parsedURL.Path, resource)

	// Finally, parsing querry params
	if len(filters) > 0 {
		q := parsedURL.Query()
		for _, fo := range filters {
			q.Set(fo.Field, fo.ToStringWF())
		}
		parsedURL.RawQuery = q.Encode()
	}

	return parsedURL.String(), nil
}

func (c *BaseClient) Close() error {
	c.HTTPClient = nil
	return nil
}
