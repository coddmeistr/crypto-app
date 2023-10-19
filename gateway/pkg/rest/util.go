package rest

import (
	"fmt"
	"net/http"
	"net/url"
)

// This function changes url of given request to a new one
// Also it copyes old url queryes
func ChangeRequestURLWithQuery(req *http.Request, newurl string) error {
	savedQuery := req.URL.Query()
	req.RequestURI = ""

	var err error
	req.URL, err = url.ParseRequestURI(newurl)
	if err != nil {
		return fmt.Errorf("failed to build URL. error: %v", err)
	}

	req.URL.RawQuery = savedQuery.Encode()
	return nil
}
