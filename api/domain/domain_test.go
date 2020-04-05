package domain

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yulpa/yulmails/api/utils"
)

type domainMock struct{}

func (e *domainMock) ListDomain() ([]*Domain, error) {
	return []*Domain{
		&Domain{
			Id:             1,
			Name:           "domain-1",
			Created:        "2019-01-25 13:34:32",
			EnvironmentID:  1,
			ConservationID: 10,
		},
	}, nil
}

func (e *domainMock) GetDomain(id int) (*Domain, error) {
	switch id {
	case 1:
		// not found
		return nil, nil
	case 2:
		return nil, errors.New("db error")
	default:
		return &Domain{
			Id:             1,
			Name:           "domain-1",
			Created:        "2019-01-25 13:34:32",
			EnvironmentID:  1,
			ConservationID: 10,
		}, nil
	}
}

func (e *domainMock) DeleteDomain(id int) error {
	switch id {
	case 1:
		return utils.NotFound
	case 2:
		return errors.New("db error")
	default:
		return nil
	}
}

func (e *domainMock) InsertDomain(a *Domain) error {
	switch a.Name {
	case "domain-1":
		return errors.New("db error")
	default:
		return nil
	}
}

func TestInsertDomain(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		environmentID  int
		conservationID int
	}{
		{
			name:           "domain-1",
			environmentID:  2,
			statusCode:     503,
			conservationID: 10,
		},
		{
			name:          "domain",
			environmentID: 2,
			statusCode:    406,
		},
		{
			name:           "domain-2",
			environmentID:  2,
			conservationID: 10,
			statusCode:     201,
		},
	}
	repo := &domainMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Insert)
	for _, test := range tests {
		payload, _ := json.Marshal(map[string]interface{}{"name": test.name, "environment_id": test.environmentID, "conservation_id": test.conservationID})
		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestDeleteDomain(t *testing.T) {
	tests := []struct {
		id         int
		statusCode int
	}{
		{
			id:         1,
			statusCode: 404,
		},
		{
			id:         2,
			statusCode: 503,
		},
		{
			id:         3,
			statusCode: 204,
		},
	}
	repo := &domainMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Delete)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodDelete, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "domain_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestGetDomain(t *testing.T) {
	tests := []struct {
		id         int
		statusCode int
	}{
		{
			id:         1,
			statusCode: 404,
		},
		{
			id:         2,
			statusCode: 503,
		},
		{
			id:         3,
			statusCode: 200,
		},
	}
	repo := &domainMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Get)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "domain_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestListDomains(t *testing.T) {
	repo := &domainMock{}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	h := &handler{repo}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.List)

	handler.ServeHTTP(rr, req)

	expected := `[{"id":1,"name":"domain-1","created":"2019-01-25 13:34:32","environment_id":1,"conservation_id":10}]`

	// we need to remove the \n since the response comes with
	res := strings.TrimSuffix(rr.Body.String(), "\n")

	assert.Equal(t, expected, res)
}
