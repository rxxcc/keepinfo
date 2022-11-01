package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID        uuid.UUID `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Password  string    `json:"-"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	// AccessLevel int `json:"access_level"`

	Contact struct {
		ID        uuid.UUID `json:"id"`
		UserID    int       `json:"user_id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Label     []string  `json:"label"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	JWT struct {
		Token string `json:"token"`
	}
)
