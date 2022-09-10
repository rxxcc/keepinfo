package router

import (
	"github.com/0xmlx/contacts-app-backend/internal/router/routes"
	"github.com/gorilla/mux"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
