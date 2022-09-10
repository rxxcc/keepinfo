package routes

import (
	"github.com/0xmlx/contacts-app-backend/internal/controllers"
	"net/http"
)

var allRoutes = []Route{

	// user routes

	{
		URI:     "/internal/auth/register",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	{
		URI:     "/internal/auth/login",
		Method:  http.MethodPost,
		Handler: controllers.Login,
	},
	{
		URI:     "/internal/auth/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		URI:     "/internal/auth/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	{
		URI:     "/internal/auth/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},

	// contacts routes

	{
		URI:     "/internal/contact",
		Method:  http.MethodGet,
		Handler: controllers.GetContacts,
	},
	{
		URI:     "/internal/contact/add",
		Method:  http.MethodPost,
		Handler: controllers.CreateContact,
	},
	{
		URI:     "/internal/contact/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetContact,
	},
	{
		URI:     "/internal/contact/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateContact,
	},
	{
		URI:     "/internal/contact/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteContact,
	},
}
