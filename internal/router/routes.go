package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inuoshios/keepinfo/internal/handlers"
	"github.com/inuoshios/keepinfo/internal/middlewares"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(middlewares.AddContentType)

	routes := r.PathPrefix("/api").Subrouter()

	routes.HandleFunc("/auth/signup", handlers.Repo.Signup).Methods(http.MethodPost)
	routes.HandleFunc("/auth/signin", handlers.Repo.Signin).Methods(http.MethodPost)
	routes.HandleFunc("/auth/signout", handlers.Repo.Signout).Methods(http.MethodPost)
	routes.HandleFunc("/auth/users", handlers.Repo.GetUsers).Methods(http.MethodGet)
	routes.HandleFunc("/auth/{id}", handlers.Repo.GetUser).Methods(http.MethodGet)
	routes.HandleFunc("/auth/{id}", handlers.Repo.DeleteUser).Methods(http.MethodDelete)
	routes.HandleFunc("/contact", handlers.Repo.GetContacts).Methods(http.MethodGet)
	routes.HandleFunc("/contact/add", handlers.Repo.CreateContact).Methods(http.MethodPost)
	routes.HandleFunc("/contact/{id}", handlers.Repo.CreateContact).Methods(http.MethodGet)
	routes.HandleFunc("/contact/{id}", handlers.Repo.UpdateContact).Methods(http.MethodPost)
	routes.HandleFunc("/contact/{id}", handlers.Repo.DeleteContact).Methods(http.MethodDelete)

	return r
}
