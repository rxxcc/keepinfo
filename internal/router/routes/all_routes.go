package routes

import (
	"net/http"

	"github.com/0xmlx/keepinfo/internal/controllers"
)

var allRoutes = []Route{

	// user routes

	{
		URI:     "/auth/signup",
		Method:  http.MethodPost,
		Handler: controllers.SignUp,
	},
	{
		URI:     "/auth/login",
		Method:  http.MethodPost,
		Handler: controllers.SignIn,
	}, {
		URI:     "/auth/logout",
		Method:  http.MethodPost,
		Handler: controllers.SignOut,
	},
	{
		URI:     "/auth/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		URI:     "/auth/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	{
		URI:     "/auth/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},

	// contacts routes

	{
		URI:     "/contact",
		Method:  http.MethodGet,
		Handler: controllers.GetContacts,
	},
	{
		URI:     "/contact/add",
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
