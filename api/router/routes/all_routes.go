package routes

import (
	"github.com/0xmlx/contacts-app-backend/api/controllers"
	"net/http"
)

var allRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	{
		URI:     "/user/create",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
	{
		URI:     "/contacts",
		Method:  http.MethodGet,
		Handler: controllers.GetContacts,
	},
	{
		URI:     "/contact/create",
		Method:  http.MethodPost,
		Handler: controllers.CreateContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteContact,
	},
}
