package abuse

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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
	GetAbuse(id int) (*abuse, error)
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

// GetAbuse returns an entity selected from the DB with its ID
func (a *abuseRepo) GetAbuse(id int) (*abuse, error) {
	query := fmt.Sprintf("SELECT name, created FROM abuse WHERE id = %d", id)
	var (
		name    string
		created string
	)
	err := a.d.QueryRow(query).Scan(&name, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.Wrapf(err, "unable to select abuse adress: %d", id)
	default:
		return &abuse{
			Name:    name,
			Created: created,
			Id:      id,
		}, nil
	}

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
// @Tags abuse
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

// Get returns to the user an abuse adress godoc
// @Summary Get an abuse adress by its ID
// @Description Returns an abuse adress by its ID
// @ID get-abuse
// @Produce  json
// @Success 200 {object} abuse
// @Success 404 {object} utils.httpError
// @Success 503 {object} utils.httpError
// @Router /abuses/{id} [get]
// @Tags abuse
// @Param id path int true "Abuse ID"
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	abuseID := ctx.Value("abuse_id").(int)
	fmt.Println(abuseID)
	a, err := h.repo.GetAbuse(abuseID)
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list abuse adresses", err.Error(),
		))
		return
	}
	if a == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	render.JSON(w, r, a)
}

func abuseCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		abuse_id, _ := strconv.Atoi(id)
		ctx := context.WithValue(r.Context(), "abuse_id", abuse_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
	r.Route("/{id}", func(r chi.Router) {
		r.Use(abuseCtx)
		r.Get("/", h.Get)
	})
	return r
}
