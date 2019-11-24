package abuse

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type AbuseRepo interface {
	ListAbuse() ([]*abuse, error)
}

type abuse struct {
	Id int `json:"id"`
	// Name is the abuse address
	Name    string `json:"name"`
	Created string `json:"created"`
}

type httpError struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *httpError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

type handler struct{ repo AbuseRepo }

// List returns to the user a list of abuse adresses godoc
// @Summary List abuses address
// @Description list abuses
// @ID list-abuses
// @Produce  json
// @Success 200 {array} abuse
// @Success 503 {object} httpError
// @Router /abuses [get]
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	e, err := h.repo.ListAbuse()
	if err != nil {
		render.Render(w, r, &httpError{
			ErrorText:      err.Error(),
			Err:            err,
			StatusText:     "unable to list abuses",
			HTTPStatusCode: http.StatusServiceUnavailable,
		})
		return
	}
	render.JSON(w, r, e)
}

type abuseRepo struct{}

// ListAbuse will return a list of abuses
func (a *abuseRepo) ListAbuse() ([]*abuse, error) {
	return []*abuse{
		&abuse{
			Id:      1,
			Name:    "abuse@local.tld",
			Created: "2019-01-25 13:34:32",
		},
	}, nil
}

// NewAbuseRepo returns a struct that implements abuse repo
// with a database connection
func NewAbuseRepo() *abuseRepo {
	return &abuseRepo{}
}

// NewRouter returns a mux router
func NewRouter(repo AbuseRepo) *chi.Mux {
	r := chi.NewRouter()
	h := &handler{repo}
	r.Get("/", h.List)
	return r
}
