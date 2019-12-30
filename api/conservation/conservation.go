package conservation

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

type handler struct{ repo ConservationRepo }

// List returns to the user a list of conservation rules godoc
// @Summary List conservation rules
// @Description list conservations
// @ID list-conservations
// @Produce  json
// @Success 200 {array} Conservation
// @Success 503 {object} utils.httpError
// @Router /conservations [get]
// @Tags conservation
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	e, err := h.repo.ListConservation()
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list conservation rules", err.Error(),
		))
		return
	}
	render.JSON(w, r, e)
}

// Get returns to the user a conservation rule godoc
// @Summary Get a conservation rule by its ID
// @Description Returns a conservation rule by its ID
// @ID get-conservation
// @Produce  json
// @Success 200 {object} Conservation
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /conservations/{id} [get]
// @Tags conservation
// @Param id path int true "Conservation ID"
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	conservationID := ctx.Value("conservation_id").(int)
	a, err := h.repo.GetConservation(conservationID)
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list conservation rules", err.Error(),
		))
		return
	}
	if a == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	render.JSON(w, r, a)
}

// Delete removes a conservation rule godoc
// @Summary Remove a conservation rule by its ID
// @Description Remove a conservation rule by its ID
// @ID delete-conservation
// @Success 204 "No Content"
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /conservations/{id} [delete]
// @Tags conservation
// @Param id path int true "Conservation ID"
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	conservationID := ctx.Value("conservation_id").(int)
	switch err := h.repo.DeleteConservation(conservationID); err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
		return
	case utils.NotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to delete conservation", err.Error(),
		))
	}
}

// Insert creates and returns a conservation rule godoc
// @Summary Insert a conservation in DB
// @Description Insert a conservation rule in DB
// @ID insert-conservation
// @Success 201 {object} Conservation
// @Success 406 "Not Acceptable. A parameter is missing"
// @Success 503 {object} utils.httpError
// @Router /conservations [post]
// @Tags conservation
// @Accept json
// @Produce  json
// @Param Conservation body conservation.Conservation true "insert conservation"
func (h *handler) Insert(w http.ResponseWriter, r *http.Request) {
	var a Conservation
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if a.Sent == 0 || a.Unsent == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	if err := h.repo.InsertConservation(&a); err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to insert conservation rule", err.Error(),
		))
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, a)
}

func conservationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		conservation_id, _ := strconv.Atoi(id)
		ctx := context.WithValue(r.Context(), "conservation_id", conservation_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewRouter returns a mux router
func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := NewConservationRepo(db)
	h := &handler{repo}
	r.Get("/", h.List)
	r.Post("/", h.Insert)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(conservationCtx)
		r.Get("/", h.Get)
		r.Delete("/", h.Delete)
	})
	return r
}
