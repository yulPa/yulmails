package abuse

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/yulpa/yulmails/api/utils"
)

type handler struct{ repo AbuseRepo }

// List returns to the user a list of abuse adresses godoc
// @Summary List abuses address
// @Description list abuses
// @ID list-abuses
// @Produce  json
// @Success 200 {array} Abuse
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
// @Success 200 {object} Abuse
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /abuses/{id} [get]
// @Tags abuse
// @Param id path int true "Abuse ID"
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	abuseID := ctx.Value("abuse_id").(int)
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

// Delete removes an abuse adress godoc
// @Summary Remove an abuse adress by its ID
// @Description Remove an abuse adress by its ID
// @ID delete-abuse
// @Success 204 "No Content"
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /abuses/{id} [delete]
// @Tags abuse
// @Param id path int true "Abuse ID"
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	abuseID := ctx.Value("abuse_id").(int)
	switch err := h.repo.DeleteAbuse(abuseID); err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
		return
	case utils.NotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to delete abuse adresses", err.Error(),
		))
	}
}

// Insert creates and returns an abuse adress godoc
// @Summary Insert an abuse address in DB
// @Description Insert an abuse address in DB
// @ID insert-abuse
// @Success 201 {object} Abuse
// @Success 406 "Not Acceptable. A parameter is missing"
// @Success 503 {object} utils.httpError
// @Router /abuses [post]
// @Tags abuse
// @Accept json
// @Produce  json
// @Param Abuse body abuse.Abuse true "insert abuse"
func (h *handler) Insert(w http.ResponseWriter, r *http.Request) {
	var a Abuse
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if len(a.Name) == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	if err := h.repo.InsertAbuse(&a); err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to insert abuse adresses", err.Error(),
		))
	}
	w.WriteHeader(http.StatusCreated)
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

// NewRouter returns a mux router
func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := NewAbuseRepo(db)
	h := &handler{repo}
	r.Get("/", h.List)
	r.Post("/", h.Insert)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(abuseCtx)
		r.Get("/", h.Get)
		r.Delete("/", h.Delete)
	})
	return r
}
