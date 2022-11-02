package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/inuoshios/keepinfo/internal/auth"
	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/repository"
	"github.com/inuoshios/keepinfo/internal/repository/dbrepo"
	"github.com/inuoshios/keepinfo/internal/response"
	"github.com/inuoshios/keepinfo/internal/utils"
)

var Repo *Repository

type Repository struct {
	App *config.Config
	DB  repository.DatabaseRepo
}

// NewRepository initializes the Repository struct
func NewRepository(a *config.Config, db *database.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(a, db.SQL),
	}
}

func NewHandlers(h *Repository) {
	Repo = h
}

// Signup - Insert user into the database
func (h *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.Email == "" {
		response.JSON(w, http.StatusUnprocessableEntity, map[string]any{"message": "email is required!"})
		return
	}
	if user.FirstName == "" {
		response.JSON(w, http.StatusUnprocessableEntity, map[string]any{"message": "firstname is required!"})
		return
	}
	if user.LastName == "" {
		response.JSON(w, http.StatusUnprocessableEntity, map[string]any{"message": "lastname is required!"})
		return
	}
	if len(user.Password) < 8 {
		response.JSON(w, http.StatusUnprocessableEntity, map[string]any{"message": "password is too short!"})
		return
	}

	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.Password = hashedPassword
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := h.DB.InsertUser(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, 200, map[string]string{
		"id": result,
	})
}

func (h *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if user.Email == "" || user.Password == "" {
		response.JSON(w, http.StatusUnprocessableEntity, map[string]any{"message": "some fields are missing"})
		return
	}

	result, err := h.DB.GetUser(user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, http.StatusNotFound, utils.ErrSqlNoRows)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.ComparePassword(result.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, fmt.Errorf("-> %w", err))
		return
	}

	token, err := auth.GenerateToken(result)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Errorf("-> %w", err))
		return
	}

	response.JSON(w, 200, models.JWT{
		Token: token,
		User:  result,
	})
}
