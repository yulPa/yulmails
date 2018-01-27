package client

import "net/http"

type SimpleHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HTTPClient struct{ SimpleHTTPClient }
