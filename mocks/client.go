package client

import (
	"github.com/yulPa/yulmails/client"

	"net/http"
)

type MockHttpClient struct{}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func NewMockClient() mongo.Session {
	return MockHttpClient{}
}
