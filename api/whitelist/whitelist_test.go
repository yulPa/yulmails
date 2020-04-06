package whitelist

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type whitelistMock struct{}

func (w *whitelistMock) ListIP() ([]string, error) {
	return []string{
		"192.168.1.10",
	}, nil
}

func TestListIP(t *testing.T) {
	repo := &whitelistMock{}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	h := &handler{repo}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.List)

	handler.ServeHTTP(rr, req)

	expected := `["192.168.1.10"]`
	// we need to remove the \n since the response comes with
	res := strings.TrimSuffix(rr.Body.String(), "\n")

	assert.Equal(t, expected, res)
}
