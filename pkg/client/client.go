package client

import "net/http"

type SimpleHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
	Close() error
}

type HTTPClient struct{ SimpleHTTPClient }

func NewHTTPClient() *HTTPClient {
	return new(HTTPClient)
}
