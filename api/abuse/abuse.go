package abuse

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/yulpa/yulmails/api/utils"
)

type abuse struct {
	Id int `json:"id"`
	// Name is the abuse address
	Name    string `json:"name"`
	Created string `json:"created"`
}

type AbuseRepo interface {
	ListAbuse() ([]*abuse, error)
}

type abuseRepo struct{}

// ListAbuse will return a list of abuses from the database
func (a *abuseRepo) ListAbuse() ([]*abuse, error) {
	return []*abuse{
		&abuse{
			Id:      1,
			Name:    "abuse@local.tld",
			Created: "2019-01-25 13:34:32",
		},
	}, nil
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
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list abuse adresses", err.Error(),
		))
		return
	}
	render.JSON(w, r, e)
}

// NewAbuseRepo returns a struct that implements abuse repo
// with a database connection
func NewAbuseRepo() *abuseRepo {
	return &abuseRepo{}
}

// NewRouter returns a mux router
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	repo := NewAbuseRepo()
	h := &handler{repo}
	r.Get("/", h.List)
	return r
}
