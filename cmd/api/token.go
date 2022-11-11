package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/inuoshios/keepinfo/internal/auth"
	"github.com/inuoshios/keepinfo/internal/models"
	"github.com/inuoshios/keepinfo/internal/response"
	"github.com/inuoshios/keepinfo/internal/utils"
)

func (h *Repository) RenewAccessToken(w http.ResponseWriter, r *http.Request) {
	var req models.RenewAccessTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	refreshPayload, err := auth.VerifyToken(req.RefreshToken)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	session, err := h.DB.GetSession(refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, http.StatusNotFound, utils.ErrSqlNoRows)
			return
		}
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if session.IsBlocked {
		response.Error(w, http.StatusUnauthorized, utils.ErrBlockedSession)
		return
	}

	if session.UserID != refreshPayload.ID.String() {
		response.Error(w, http.StatusUnauthorized, utils.ErrIncorrectSessionUser)
		return
	}

	if session.RefreshToken != req.RefreshToken {
		response.Error(w, http.StatusUnauthorized, utils.ErrMismatchedToken)
		return
	}

	if time.Now().After(session.ExpiredAt) {
		response.Error(w, http.StatusUnauthorized, utils.ErrExpiredSession)
		return
	}

	acessToken, accessPayload, err := auth.GenerateToken(session.ID, time.Duration(time.Minute*4))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, fmt.Errorf("-> %w", err))
		return
	}

	response.JSON(w, http.StatusOK, models.RenewAccessTokenResponse{
		AccessToken:          acessToken,
		AccessTokenExpiresAt: accessPayload.ExpiresAt.Time,
	})
}
