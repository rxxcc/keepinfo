package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/repository"
	"github.com/inuoshios/keepinfo/internal/response"
	v "github.com/inuoshios/keepinfo/internal/validator"
)

type createUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type Handler struct {
	appRepo repository.Repository
}

func _(db *database.DB) *Handler {
	return &Handler{
		appRepo: repository.NewDatabase(db.SQL),
	}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var user createUser

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

	args := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	result, err := h.appRepo.InsertUser(args)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, 200, result)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "login")
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "logout")
}

// GetUsers should be a protected route
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get users")
}

// GetUser should be a protected route
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get user")
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete user")
}

func (h *Handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create contact")
}

func (h *Handler) GetContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get contacts")
}

func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get contact")
}

func (h *Handler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update contact")
}

func (h *Handler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete contact")
}
