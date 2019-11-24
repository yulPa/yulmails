package abuse

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

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

type abuseRepo struct{ d *sql.DB }

// ListAbuse will return a list of abuses from the database
func (a *abuseRepo) ListAbuse() ([]*abuse, error) {
	query := "SELECT * from abuse;"
	abuses := make([]*abuse, 0)
	res, err := a.d.Query(query)
	if err != nil {
		return abuses, errors.Wrapf(err, "unable to query db: %s", query)
	}
	defer res.Close()
	for res.Next() {
		var (
			id      int
			name    string
			created string
		)
		if err := res.Scan(&id, &name, &created); err != nil {
			return abuses, errors.Wrap(err, "unable to extract result")
		}
		abuses = append(abuses, &abuse{
			Id:      id,
			Name:    name,
			Created: created,
		})
	}
	return abuses, nil
}

type handler struct{ repo AbuseRepo }

// List returns to the user a list of abuse adresses godoc
// @Summary List abuses address
// @Description list abuses
// @ID list-abuses
// @Produce  json
// @Success 200 {array} abuse
// @Success 503 {object} utils.httpError
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
func NewAbuseRepo(db *sql.DB) *abuseRepo {
	return &abuseRepo{d: db}
}

// NewRouter returns a mux router
func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := NewAbuseRepo(db)
	h := &handler{repo}
	r.Get("/", h.List)
	return r
}
