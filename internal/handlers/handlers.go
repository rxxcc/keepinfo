package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/inuoshios/keepinfo/internal/database"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/response"
	v "github.com/inuoshios/keepinfo/internal/validator"
)

type Handler struct {
	*database.DB
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := h.InsertUser(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, 200, result)
}
