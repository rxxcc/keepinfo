package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/inuoshios/keepinfo/internal/models"
)

// payload
type Claims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

var user models.User

func NewClaims(username string, duration time.Duration) (*Claims, error) {
	randomID, _ := uuid.NewRandom()
	claims := Claims{
		ID:       randomID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.Email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	return &claims, nil
}
