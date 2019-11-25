package abuse

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
		return &abuse{
			Id:      1,
			Name:    "abuse-1",
			Created: "2019-01-25 13:34:32",
		}, nil
	case 3:
		return nil, errors.New("db error")
	default:
		// not found
		return nil, nil
	}
}

func TestGetAbuse(t *testing.T) {
	tests := []struct {
		id         int
		statusCode int
	}{
		{
			id:         1,
			statusCode: 200,
		},
		{
			id:         2,
			statusCode: 404,
		},
		{
			id:         3,
			statusCode: 503,
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
