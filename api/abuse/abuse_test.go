package abuse

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

type abuseMock struct{}

func (e *abuseMock) ListAbuse() ([]*abuse, error) {
	return []*abuse{
		&abuse{
			Id:      1,
			Name:    "abuse-1",
			Created: "2019-01-25 13:34:32",
		},
	}, nil
}

func (e *abuseMock) GetAbuse(id int) (*abuse, error) {
	switch id {
	case 1:
		// not found
		return nil, nil
	case 2:
		return nil, errors.New("db error")
	default:
		return &abuse{
			Id:      1,
			Name:    "abuse-1",
			Created: "2019-01-25 13:34:32",
		}, nil
	}
}

func (e *abuseMock) DeleteAbuse(id int) error {
	switch id {
	case 1:
		return utils.NotFound
	case 2:
		return errors.New("db error")
	default:
		return nil
	}
}

func (e *abuseMock) InsertAbuse(a *abuse) error {
	switch a.Name {
	case "abuse-1":
		return errors.New("db error")
	default:
		return nil
	}
}

func TestInsertAbuse(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "abuse-1",
			statusCode: 503,
		},
		{
			name:       "",
			statusCode: 406,
		},
		{
			name:       "abuse-2",
			statusCode: 201,
		},
	}
	repo := &abuseMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Insert)
	for _, test := range tests {
		payload, _ := json.Marshal(map[string]string{"name": test.name})
		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestDeleteAbuse(t *testing.T) {
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
	repo := &abuseMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Delete)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodDelete, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "abuse_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestGetAbuse(t *testing.T) {
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
	repo := &abuseMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Get)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "abuse_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestListAbuses(t *testing.T) {
	repo := &abuseMock{}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	h := &handler{repo}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.List)

	handler.ServeHTTP(rr, req)

	expected := `[{"id":1,"name":"abuse-1","created":"2019-01-25 13:34:32"}]`
	// we need to remove the \n since the response comes with
	res := strings.TrimSuffix(rr.Body.String(), "\n")

	assert.Equal(t, expected, res)
}
