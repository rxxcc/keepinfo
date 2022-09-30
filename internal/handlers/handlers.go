package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inuoshios/keepinfo/internal/data"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/response"
	v "github.com/inuoshios/keepinfo/internal/validator"
)

type Handlers struct {
	Data data.Database
}

func (h *Handlers) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	hashedPassword, err := v.Hash(user.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = hashedPassword

	// err = h.Data.InsertUser(user)
	// if err != nil {
	// 	response.Error(w, http.StatusInternalServerError, err)
	// 	return
	// }

	response.JSON(w, 200, user)
}

func (h *Handlers) Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "signin")
}

func (h *Handlers) Signout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "signout")
}

// should be a protected route
func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get users")
}

// should be a protected route
func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get user")
}

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete user")
}

func (h *Handlers) CreateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create contact")
}

func (h *Handlers) GetContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get contacts")
}

func (h *Handlers) GetContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get contact")
}

func (h *Handlers) UpdateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update contact")
}

func (h *Handlers) DeleteContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete contact")
}
