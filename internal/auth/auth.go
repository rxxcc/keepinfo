package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/inuoshios/keepinfo/internal/utils"
)

func GenerateToken(id uuid.UUID) (string, *Claims, error) {
	claims, err := NewClaims(id)
	if err != nil {
		return "", claims, fmt.Errorf("payload: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", claims, fmt.Errorf("error creating signed string: %w", err)
	}

	return signedStr, claims, nil
}

// VerifyToken checks if the token is valid or not
func VerifyToken(token string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, utils.ErrInvalidToken
		}
		return []byte(os.Getenv("SECRET")), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, utils.ErrExpiredToken) {
			return nil, utils.ErrExpiredToken
		}
	}

	payload, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, utils.ErrInvalidToken
	}

	return payload, nil
}
