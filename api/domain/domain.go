package domain

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

type handler struct{ repo DomainRepo }

// List returns to the user a list of domain godoc
// @Summary List domains
// @Description list domains
// @ID list-domains
// @Produce  json
// @Success 200 {array} Domain
// @Success 503 {object} utils.httpError
// @Router /domains [get]
// @Tags domain
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	e, err := h.repo.ListDomain()
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list domain", err.Error(),
		))
		return
	}
	render.JSON(w, r, e)
}

// Get returns to the user an domain godoc
// @Summary Get an domain by its ID
// @Description Returns an domain by its ID
// @ID get-domain
// @Produce  json
// @Success 200 {object} Domain
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /domains/{id} [get]
// @Tags domain
// @Param id path int true "Domain ID"
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	domainID := ctx.Value("domain_id").(int)
	a, err := h.repo.GetDomain(domainID)
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list domain", err.Error(),
		))
		return
	}
	if a == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	render.JSON(w, r, a)
}

// Delete removes an domain adress godoc
// @Summary Remove an domain adress by its ID
// @Description Remove an domain adress by its ID
// @ID delete-domain
// @Success 204 "No Content"
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /domains/{id} [delete]
// @Tags domain
// @Param id path int true "Domain ID"
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	domainID := ctx.Value("domain_id").(int)
	switch err := h.repo.DeleteDomain(domainID); err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
		return
	case utils.NotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to delete domain", err.Error(),
		))
	}
}

// Insert creates and returns an domain godoc
// @Summary Insert an domain in DB
// @Description Insert an domain in DB
// @ID insert-domain
// @Success 201 {object} Domain
// @Success 406 "Not Acceptable. A parameter is missing"
// @Success 503 {object} utils.httpError
// @Router /domains [post]
// @Tags domain
// @Accept json
// @Produce  json
// @Param Domain body domain.Domain true "insert entity"
func (h *handler) Insert(w http.ResponseWriter, r *http.Request) {
	var a Domain
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if len(a.Name) == 0 || a.ConservationID == 0 || a.EnvironmentID == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	if err := h.repo.InsertDomain(&a); err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to insert domain", err.Error(),
		))
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, a)
}

func domainCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		domain_id, _ := strconv.Atoi(id)
		ctx := context.WithValue(r.Context(), "domain_id", domain_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewRouter returns a mux router
func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := NewDomainRepo(db)
	h := &handler{repo}
	r.Get("/", h.List)
	r.Post("/", h.Insert)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(domainCtx)
		r.Get("/", h.Get)
		r.Delete("/", h.Delete)
	})
	return r
}
