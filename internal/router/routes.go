package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inuoshios/keepinfo/internal/handlers"
	"github.com/inuoshios/keepinfo/internal/middleware"
)

var routes *handlers.Handler

func NEW() http.Handler {
	mux := chi.NewRouter()

	// path prefix using chi
	mux.Route("/api", func(mux chi.Router) {
		// adding the `ContentType` middleware
		mux.Use(middleware.AddContentType)
		// adding the routes
		mux.Post("/auth/signup", routes.Signup)
	})

	return mux
}
