package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) http.Handler {
	api := &API{db: db}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", api.health)
	mux.HandleFunc("GET /api/v1/dashboard", api.dashboard)
	mux.HandleFunc("GET /api/v1/contacts", api.contacts)

	return logging(mux)
}

func (a *API) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func (a *API) dashboard(w http.ResponseWriter, r *http.Request) {
	// MVP placeholder: organization scoping will come from authenticated user context.
	writeJSON(w, http.StatusOK, map[string]any{
		"new_leads":     0,
		"appointments":  0,
		"follow_ups":    0,
		"applications":  0,
		"underwriting":  0,
		"message":       "Dashboard endpoint ready for authenticated organization scope.",
	})
}

func (a *API) contacts(w http.ResponseWriter, r *http.Request) {
	// MVP placeholder until auth middleware supplies organization_id.
	writeJSON(w, http.StatusOK, []any{})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
