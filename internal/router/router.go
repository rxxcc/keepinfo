package router

import (
	"github.com/gorilla/mux"
	"github.com/inuoshios/keepinfo/internal/router/routes"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
