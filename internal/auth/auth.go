package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/inuoshios/keepinfo/internal/models"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"jti": user.ID,
		"iss": user.Email,
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	signedStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", fmt.Errorf("error creating signed string: %w", err)
	}

	return signedStr, nil
}
