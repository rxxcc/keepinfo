package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/data"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/response"
	v "github.com/inuoshios/keepinfo/internal/validator"
)

var (
	user      models.User
	Repo      *Repository
	Validator v.Validator
	Data      *data.DB
)

type Repository struct {
	App *config.Config
}

func (repo *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)

	existingUser, err := Data.GetUserByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	// validate email
	Validator.CheckField(user.Email == "", "Email", "Email address is required!")
	Validator.CheckField(v.Matches(user.Email, v.RgxEmail), "Email", "Must be a valid email address.")
	Validator.CheckField(existingUser != nil, "Email", "Email is already in use.")

	// validate password
	Validator.CheckField(user.Password != "", "Password", "Password is required!")
	Validator.CheckField(len(user.Password) < 8, "Password", "Password is too short.")
	Validator.CheckField(len(user.Password) > 72, "Password", "Password is too long.")

	hashedPassword, err := v.Hash(user.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, 200, user)

	fmt.Println(hashedPassword)
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
