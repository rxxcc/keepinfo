package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/response"
)

func (h *Repository) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	contact.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	contact.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := h.DB.InsertContact(&contact)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, 200, map[string]string{
		"id": result,
	})
}
