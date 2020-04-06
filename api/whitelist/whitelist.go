package whitelist

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/yulpa/yulmails/api/utils"
)

type handler struct{ repo WhitelistRepo }

// List returns to the user a list of whitelisted IP godoc
// @Summary List whitelisted IP address
// @Description list whitelisted IP address
// @ID list-whitelistd-ips
// @Produce  json
// @Success 200 {array} string
// @Success 503 {object} utils.httpError
// @Router /whitelist [get]
// @Tags whitelist
func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	i, err := h.repo.ListIP()
	if err != nil {
		render.Render(w, r, utils.NewHTTPError(
			err, http.StatusServiceUnavailable, "unable to list abuse adresses", err.Error(),
		))
		return
	}
	render.JSON(w, r, i)
}

// NewRouter returns a mux router
func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := NewWhitelistRepo(db)
	h := &handler{repo}
	r.Get("/", h.List)
	return r
}
