package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	User struct {
		Username  string    `json:"username"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	// AccessLevel int `json:"access_level"`

	Contact struct {
		ID        uuid.UUID      `json:"id"`
		UserID    string         `json:"user_id"`
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"email"`
		Phone     string         `json:"phone"`
		Label     pq.StringArray `json:"label"`
		Address   string         `json:"address"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt time.Time      `json:"deleted_at"`
	}

	JWT struct {
		SessionID             uuid.UUID `json:"session_id"`
		AccessToken           string    `json:"access_token"`
		RefreshToken          string    `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
		AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
		User                  User      `json:"user"`
	}

	Session struct {
		ID           uuid.UUID `json:"id"`
		UserID       string    `json:"user_id"`
		RefreshToken string    `json:"refresh_token"`
		UserAgent    string    `json:"user_agent"`
		ClientIP     string    `json:"client_ip"`
		IsBlocked    bool      `json:"is_blocked"`
		ExpiredAt    time.Time `json:"expired_at"`
		CreatedAt    time.Time `json:"created_at"`
	}

	GetAllUsers struct {
		UserID string `json:"user_id"`
	}

	RenewAccessTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	RenewAccessTokenResponse struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	}
)
