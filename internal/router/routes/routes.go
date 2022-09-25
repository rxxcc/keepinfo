package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inuoshios/keepinfo/internal/middlewares"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := allRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		// Add middlewares
		r.Use(middlewares.AddContentType)

		routes := r.PathPrefix("/api").Subrouter()
		routes.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
