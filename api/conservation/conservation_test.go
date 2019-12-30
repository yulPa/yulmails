package conservation

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

type conservationMock struct{}

func (e *conservationMock) ListConservation() ([]*Conservation, error) {
	return []*Conservation{
		&Conservation{
			ID:               1,
			Created:          "2019-01-25 13:34:32",
			Sent:             10,
			Unsent:           5,
			KeepEmailContent: true,
		},
	}, nil
}

func (e *conservationMock) GetConservation(id int) (*Conservation, error) {
	switch id {
	case 1:
		// not found
		return nil, nil
	case 2:
		return nil, errors.New("db error")
	default:
		return &Conservation{
			ID:      1,
			Created: "2019-01-25 13:34:32",
			Sent:    10,
			Unsent:  5,
		}, nil
	}
}

func (e *conservationMock) DeleteConservation(id int) error {
	switch id {
	case 1:
		return utils.NotFound
	case 2:
		return errors.New("db error")
	default:
		return nil
	}
}

func (e *conservationMock) InsertConservation(a *Conservation) error {
	switch a.Sent {
	case 17:
		return errors.New("db error")
	default:
		return nil
	}
}

func TestInsertConservation(t *testing.T) {
	tests := []struct {
		statusCode       int
		sent             int
		unsent           int
		keepEmailContent bool
	}{
		{
			sent:             17,
			unsent:           5,
			keepEmailContent: false,
			statusCode:       503,
		},
		{
			unsent:           5,
			sent:             0,
			keepEmailContent: true,
			statusCode:       406,
		},
		{
			sent:             10,
			unsent:           5,
			keepEmailContent: true,
			statusCode:       201,
		},
	}
	repo := &conservationMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Insert)
	for _, test := range tests {
		payload, _ := json.Marshal(map[string]interface{}{"sent": test.sent, "unsent": test.unsent, "keep_email_content": test.keepEmailContent})
		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestDeleteConservation(t *testing.T) {
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
	repo := &conservationMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Delete)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodDelete, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "conservation_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestGetConservation(t *testing.T) {
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
	repo := &conservationMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Get)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "conservation_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestListConservations(t *testing.T) {
	repo := &conservationMock{}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	h := &handler{repo}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.List)

	handler.ServeHTTP(rr, req)

	expected := `[{"id":1,"created":"2019-01-25 13:34:32","sent":10,"unsent":5,"keep_email_content":true}]`
	// we need to remove the \n since the response comes with
	res := strings.TrimSuffix(rr.Body.String(), "\n")

	assert.Equal(t, expected, res)
}
