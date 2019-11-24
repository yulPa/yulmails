package entity

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/yulpa/yulmails/api/utils"
)

type EntityRepo interface {
	ListEntity() ([]*entity, error)
}

type entity struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Created     string `json:"created"`
	Description string `json:"description"`
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

type handler struct{ repo EntityRepo }

// List returns to the user a list of enties godoc
// @Summary List entities
// @Description list entities
// @ID list-entities
// @Produce  json
// @Success 200 {array} entity
// @Success 503 {object} utils.httpError
// @Router /entities [get]
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	e, err := h.repo.ListEntity()
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list entities", err.Error(),
		))
		return
	}
	render.JSON(w, r, e)
}

type entityRepo struct{}

// ListEntity will return a list of entities
func (e *entityRepo) ListEntity() ([]*entity, error) {
	return []*entity{
		&entity{
			Id:          1,
			Name:        "entity-1",
			Created:     "2019-01-25 13:34:32",
			Description: "the first entity",
		},
	}, nil
}

// NewEntityRepo returns a struct that implements entity repo
// with a database connection
func NewEntityRepo() *entityRepo {
	return &entityRepo{}
}

// NewRouter returns a mux router
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	repo := NewEntityRepo()
	h := &handler{repo}
	r.Get("/", h.List)
	return r
}
