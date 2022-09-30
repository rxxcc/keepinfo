package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inuoshios/keepinfo/internal/handlers"
	"github.com/inuoshios/keepinfo/internal/middlewares"
)

var Repo *handlers.Handlers

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(middlewares.AddContentType)

	routes := r.PathPrefix("/api").Subrouter()

	routes.HandleFunc("/auth/signup", Repo.Signup).Methods(http.MethodPost)
	routes.HandleFunc("/auth/signin", Repo.Signin).Methods(http.MethodPost)
	routes.HandleFunc("/auth/signout", Repo.Signout).Methods(http.MethodPost)
	routes.HandleFunc("/auth/users", Repo.GetUsers).Methods(http.MethodGet)
	routes.HandleFunc("/auth/{id}", Repo.GetUser).Methods(http.MethodGet)
	routes.HandleFunc("/auth/{id}", Repo.DeleteUser).Methods(http.MethodDelete)
	routes.HandleFunc("/contact", Repo.GetContacts).Methods(http.MethodGet)
	routes.HandleFunc("/contact/add", Repo.CreateContact).Methods(http.MethodPost)
	routes.HandleFunc("/contact/{id}", Repo.CreateContact).Methods(http.MethodGet)
	routes.HandleFunc("/contact/{id}", Repo.UpdateContact).Methods(http.MethodPost)
	routes.HandleFunc("/contact/{id}", Repo.DeleteContact).Methods(http.MethodDelete)

	return r
}
