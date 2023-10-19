package rest

import (
	"io"
	"net/http"
	"net/url"
)

type APIResponse struct {
	response *http.Response
}

func (ar *APIResponse) Body() io.ReadCloser {
	return ar.response.Body
}

func (ar *APIResponse) Response() *http.Response {
	return ar.response
}

func (ar *APIResponse) ReadBody() ([]byte, error) {
	defer ar.response.Body.Close()
	return io.ReadAll(ar.response.Body)
}

func (ar *APIResponse) StatusCode() int {
	return ar.response.StatusCode
}

func (ar *APIResponse) Location() (*url.URL, error) {
	return ar.response.Location()
}
