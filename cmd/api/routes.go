package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inuoshios/keepinfo/internal/handlers"
)

func NEW() http.Handler {

	mux := chi.NewRouter()

	// Adding the middlewares
	mux.Use(AddContentType)

	// Paths
	mux.Route("/auth", func(r chi.Router) {
		r.Post("/signup", handlers.Repo.Signup)
	})

	mux.Route("/api", func(r chi.Router) {
		r.Post("/contacts", handlers.Repo.CreateContact)
	})

	return mux
}
