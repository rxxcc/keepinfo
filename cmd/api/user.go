package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/inuoshios/keepinfo/internal/auth"
	"github.com/inuoshios/keepinfo/internal/config"
	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/repository"
	"github.com/inuoshios/keepinfo/internal/repository/postgres"
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
		DB:  postgres.NewPostgresRepo(a, db.SQL),
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

	if err := utils.ValidateEmail(user.Email); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	names := []string{user.FirstName, user.LastName}
	if err := utils.ValidateName(names...); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
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
		"status": "success",
		"id":     result,
	})
}

func (h *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.ValidateEmail(user.Email); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
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

	if err = utils.ComparePassword(result.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, fmt.Errorf("-> %w", err))
		return
	}

	accessToken, accessPayload, err := auth.GenerateToken(result.ID, time.Duration(time.Minute*1))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Errorf("-> %w", err))
		return
	}

	refreshToken, refreshPayload, err := auth.GenerateToken(result.ID, time.Duration(time.Minute*4))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Errorf("-> %w", err))
		return
	}

	randomID, _ := uuid.NewRandom()

	session, err := h.DB.CreateSession(&models.Session{
		ID:           randomID,
		UserID:       result.ID.String(),
		RefreshToken: refreshToken,
		UserAgent:    r.UserAgent(),
		ClientIP:     r.RemoteAddr,
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiresAt.Time,
	})

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, 200, models.JWT{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt.Time,
		User:                  result,
	})
}
