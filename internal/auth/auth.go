package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/inuoshios/keepinfo/internal/models"
)

func GenerateToken(user models.User) (string, error) {
	tokenID, _ := uuid.NewRandom()

	claims := struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", fmt.Errorf("error creating signed string: %w", err)
	}

	return signedStr, nil
}
