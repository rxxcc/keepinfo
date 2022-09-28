package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/responses"
)

var (
	user models.User
	Repo *Repository
)

type Repository struct {
	App *config.Config
}

func (repo *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)

	responses.JSON(w, 200, user)
}

func (repo *Repository) Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "signin")
}

func (repo *Repository) Signout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "signout")
}

// should be a protected route
func (repo *Repository) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get users")
}

// should be a protected route
func (repo *Repository) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get user")
}

func (repo *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete user")
}

func (repo *Repository) CreateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create contact")
}

func (repo *Repository) GetContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get contacts")
}

func (repo *Repository) GetContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get contact")
}

func (repo *Repository) UpdateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update contact")
}

func (repo *Repository) DeleteContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete contact")
}
