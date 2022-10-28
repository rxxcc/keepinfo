package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inuoshios/keepinfo/internal/handlers"
	"github.com/inuoshios/keepinfo/internal/middleware"
)

func NEW() http.Handler {

	mux := chi.NewRouter()

	// Adding the middlewares
	mux.Use(middleware.AddContentType)

	// Paths
	mux.Route("/api", func(r chi.Router) {
		r.Post("/auth/signup", handlers.Repo.Signup)
	})

	return mux
}
