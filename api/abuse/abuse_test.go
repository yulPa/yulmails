package abuse

import (
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

func TestGetAbuses(t *testing.T) {
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
