package entity

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

type entityMock struct{}

func (e *entityMock) ListEntity() ([]*Entity, error) {
	return []*Entity{
		&Entity{
			Id:             1,
			Name:           "entity-1",
			Created:        "2019-01-25 13:34:32",
			Description:    "description-1",
			ConservationID: 1,
			AbuseID:        2,
		},
	}, nil
}

func (e *entityMock) GetEntity(id int) (*Entity, error) {
	switch id {
	case 1:
		// not found
		return nil, nil
	case 2:
		return nil, errors.New("db error")
	default:
		return &Entity{
			Id:             1,
			Name:           "entity-1",
			Created:        "2019-01-25 13:34:32",
			Description:    "description-1",
			ConservationID: 1,
			AbuseID:        2,
		}, nil
	}
}

func (e *entityMock) DeleteEntity(id int) error {
	switch id {
	case 1:
		return utils.NotFound
	case 2:
		return errors.New("db error")
	default:
		return nil
	}
}

func (e *entityMock) InsertEntity(a *Entity) error {
	switch a.Name {
	case "entity-1":
		return errors.New("db error")
	default:
		return nil
	}
}

func TestInsertEntity(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		abuseID        int
		conservationID int
		description    string
	}{
		{
			name:           "entity-1",
			description:    "description-1",
			abuseID:        1,
			conservationID: 2,
			statusCode:     503,
		},
		{
			name:           "entity",
			description:    "description",
			conservationID: 2,
			abuseID:        0,
			statusCode:     406,
		},
		{
			name:           "entity-2",
			description:    "description-2",
			abuseID:        1,
			conservationID: 2,
			statusCode:     201,
		},
	}
	repo := &entityMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Insert)
	for _, test := range tests {
		payload, _ := json.Marshal(map[string]interface{}{"name": test.name, "description": test.description, "abuse_id": test.abuseID, "conservation_id": test.conservationID})
		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestDeleteEntity(t *testing.T) {
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
	repo := &entityMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Delete)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodDelete, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "entity_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestGetEntity(t *testing.T) {
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
	repo := &entityMock{}
	h := &handler{repo}
	handler := http.HandlerFunc(h.Get)
	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()
		// we set manually the value in the context
		ctx := context.WithValue(req.Context(), "entity_id", test.id)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, test.statusCode, rr.Result().StatusCode)
	}
}

func TestListEntities(t *testing.T) {
	repo := &entityMock{}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	h := &handler{repo}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.List)

	handler.ServeHTTP(rr, req)

	expected := `[{"id":1,"name":"entity-1","created":"2019-01-25 13:34:32","description":"description-1","conservation_id":1,"abuse_id":2}]`
	// we need to remove the \n since the response comes with
	res := strings.TrimSuffix(rr.Body.String(), "\n")

	assert.Equal(t, expected, res)
}
