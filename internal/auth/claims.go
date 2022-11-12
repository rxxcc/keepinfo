package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// payload
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewClaims(username string, duration time.Duration) (*Claims, error) {
	randomID, _ := uuid.NewRandom()
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        randomID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	return &claims, nil
}
