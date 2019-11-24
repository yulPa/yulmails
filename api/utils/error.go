package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

type httpError struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	ErrorText      string `json:"error,omitempty"`
}

// NewHTTPError creates and returns a new httpError.
// err is the low-level runtime error, statusCode is the HTTP response status code,
// statusText stands for the user-level status message,
// finally errorText is the application-level error message (for debugging)
func NewHTTPError(err error, statusCode int, statusText string, errorText string) *httpError {
	return &httpError{
		Err:            err,
		HTTPStatusCode: statusCode,
		StatusText:     statusText,
		ErrorText:      errorText,
	}
}

// Render renders a single payload and respond to the client request
func (e *httpError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
