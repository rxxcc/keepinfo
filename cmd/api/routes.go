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
		r.Post("/signin", handlers.Repo.Login)
	})

	// contact url path
	mux.Group(func(r chi.Router) {
		r.Use(VerifyToken)

		r.Route("/api", func(c chi.Router) {
			c.Post("/contacts", handlers.Repo.CreateContact)
			c.Get("/contacts", handlers.Repo.GetContacts)
		})
	})

	return mux
}
