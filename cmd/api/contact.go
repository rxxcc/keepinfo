package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/inuoshios/keepinfo/internal/auth"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/response"
	"github.com/inuoshios/keepinfo/internal/utils"
)

func (h *Repository) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	authPayload := r.Context().Value(authPayloadKey).(*auth.Claims)

	createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	var args = models.Contact{
		ID:        contact.ID,
		UserID:    authPayload.ID.String(),
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		Label:     contact.Label,
		Address:   contact.Address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: contact.DeletedAt,
	}

	result, err := h.DB.InsertContact(&args)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, 200, map[string]string{
		"id": result,
	})
}

func (h *Repository) GetContacts(w http.ResponseWriter, r *http.Request) {
	// var cnt models.Contact
	authPayload := r.Context().Value(authPayloadKey).(*auth.Claims)

	args := models.GetAllUsers{
		UserID: authPayload.ID.String(),
	}

	result, err := h.DB.GetContacts(args)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Errorf("-> %w", err))
		return
	}

	response.JSON(w, 200, result)

}

func (h *Repository) GetContact(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "id")

	result, err := h.DB.GetContact(path)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, http.StatusNotFound, utils.ErrSqlNoRows)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	authPayload := r.Context().Value(authPayloadKey).(*auth.Claims)
	if result.UserID != authPayload.ID.String() {
		response.Error(w, http.StatusUnauthorized, utils.ErrAuthUser)
		return
	}

	response.JSON(w, http.StatusOK, result)
}