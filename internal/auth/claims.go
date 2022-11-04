package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/inuoshios/keepinfo/internal/models"
)

type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

var user models.User

func NewClaims(id uuid.UUID) (*Claims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			Issuer:    user.Email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
	}
	return &claims, nil
}
