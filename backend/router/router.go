// Registers all backend HTTP and WebSocket routes.
package router

import (
	"encoding/json"
	"net/http"

	"amber/backend/environment"
	"amber/backend/era"
	"amber/backend/health"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func New() http.Handler {
	r := chi.NewRouter()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	environmentHandler := environment.NewHandler(environment.NewService())

	r.Get("/health", health.Handler)
	r.Get("/eras", listEras)
	r.Get("/environment/{era}/{lat}/{lng}", environmentHandler.PreviewByPath)
	r.Post("/environment/preview", environmentHandler.Preview)

	return r
}

func listEras(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(era.List())
}
