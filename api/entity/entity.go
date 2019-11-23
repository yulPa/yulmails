package entity

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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

type handler struct{ repo EntityRepo }

// List writes the list of entities in the response writer
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	e, err := h.repo.ListEntity()
	fmt.Println(len(e))
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to list entities: %s", err), http.StatusServiceUnavailable)
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
