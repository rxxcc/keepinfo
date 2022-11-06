package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NEW() http.Handler {

	mux := chi.NewRouter()

	// Adding the middlewares
	mux.Use(AddContentType)

	// signup/signin paths
	mux.Route("/auth", func(r chi.Router) {
		r.Post("/signup", Repo.Signup)
		r.Post("/signin", Repo.Login)
	})

	// contact url path
	mux.Group(func(r chi.Router) {
		r.Use(VerifyToken)

		r.Route("/api", func(c chi.Router) {
			c.Post("/contacts", Repo.CreateContact)
			c.Get("/contacts", Repo.GetContacts)
			c.Get("/contacts/{id}", Repo.GetContact)
			c.Patch("/contacts/{id}", Repo.UpdateContact)
			c.Delete("/contacts/{id}", Repo.DeleteContact)
		})
	})

	return mux
}
