package mocks

import (
	"net/http"

	"github.com/yulPa/yulmails/pkg/client"
)

type MockHttpClient struct{}

func (m MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func (m MockHttpClient) Close() error {
	return nil
}

func NewMockClient() client.SimpleHTTPClient {
	return MockHttpClient{}
}
