package entity

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

type handler struct{ repo EntityRepo }

// List returns to the user a list of entity godoc
// @Summary List entities
// @Description list entities
// @ID list-entities
// @Produce  json
// @Success 200 {array} Entity
// @Success 503 {object} utils.httpError
// @Router /entities [get]
// @Tags entity
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	e, err := h.repo.ListEntity()
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list entity", err.Error(),
		))
		return
	}
	render.JSON(w, r, e)
}

// Get returns to the user an entity godoc
// @Summary Get an entity by its ID
// @Description Returns an entity by its ID
// @ID get-entity
// @Produce  json
// @Success 200 {object} Entity
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /entities/{id} [get]
// @Tags entity
// @Param id path int true "Entity ID"
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entityID := ctx.Value("entity_id").(int)
	a, err := h.repo.GetEntity(entityID)
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list entity", err.Error(),
		))
		return
	}
	if a == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	render.JSON(w, r, a)
}

// Delete removes an entity adress godoc
// @Summary Remove an entity adress by its ID
// @Description Remove an entity adress by its ID
// @ID delete-entity
// @Success 204 "No Content"
// @Success 404 "Not Found"
// @Success 503 {object} utils.httpError
// @Router /entities/{id} [delete]
// @Tags entity
// @Param id path int true "Entity ID"
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entityID := ctx.Value("entity_id").(int)
	switch err := h.repo.DeleteEntity(entityID); err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
		return
	case utils.NotFound:
		w.WriteHeader(http.StatusNotFound)
		return
	default:
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to delete entity", err.Error(),
		))
	}
}

// Insert creates and returns an entity godoc
// @Summary Insert an entity in DB
// @Description Insert an entity in DB
// @ID insert-entity
// @Success 201 {object} Entity
// @Success 406 "Not Acceptable. A parameter is missing"
// @Success 503 {object} utils.httpError
// @Router /entities [post]
// @Tags entity
// @Accept json
// @Produce  json
// @Param Entity body entity.Entity true "insert entity"
func (h *handler) Insert(w http.ResponseWriter, r *http.Request) {
	var a Entity
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if len(a.Name) == 0 || len(a.Description) == 0 || a.ConservationID == 0 || a.AbuseID == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	if err := h.repo.InsertEntity(&a); err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to insert entity", err.Error(),
		))
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, a)
}

func entityCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		entity_id, _ := strconv.Atoi(id)
		ctx := context.WithValue(r.Context(), "entity_id", entity_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewRouter returns a mux router
func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := NewEntityRepo(db)
	h := &handler{repo}
	r.Get("/", h.List)
	r.Post("/", h.Insert)
	r.Route("/{id}", func(r chi.Router) {
		r.Use(entityCtx)
		r.Get("/", h.Get)
		r.Delete("/", h.Delete)
	})
	return r
}
