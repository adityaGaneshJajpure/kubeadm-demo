package httpclient

import (
	"net/http"
)

type IHTTPClient interface {
	Get(url string) (*http.Response, error)
}

type HTTPClient struct{}

func (c *HTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
