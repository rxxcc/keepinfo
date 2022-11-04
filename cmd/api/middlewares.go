package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/inuoshios/keepinfo/internal/response"
	"github.com/inuoshios/keepinfo/internal/utils"
)

// AddContentType functions helps set headers for our api.
func AddContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Max-Age", "86400")

			next.ServeHTTP(w, r)
		})
}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			authTypeBearer := "Bearer"
			keyfunc := func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET")), nil
			}

			if len(authHeader) == 0 {
				response.Error(w, http.StatusUnauthorized, utils.ErrAuthHeader)
				return
			}

			bearerToken := strings.Split(authHeader, " ")

			if len(bearerToken) < 2 {
				response.Error(w, http.StatusUnauthorized, utils.ErrInvalidAuthHeader)
				return
			}

			authType := bearerToken[0]
			if authType != authTypeBearer {
				response.Error(w, http.StatusUnauthorized, fmt.Errorf("%w, %s", utils.ErrUnsupportedAuthType, authType))
				return
			}

			accessToken := bearerToken[1]

			token, err := jwt.Parse(accessToken, keyfunc)
			if err != nil {
				verr, ok := err.(*jwt.ValidationError)
				if ok && errors.Is(verr.Inner, utils.ErrExpiredToken) {
					response.Error(w, http.StatusUnauthorized, utils.ErrExpiredToken)
					return
				}
				response.Error(w, http.StatusUnauthorized, utils.ErrInvalidToken)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				response.JSON(w, http.StatusAccepted, claims)
				next.ServeHTTP(w, r)
			}
		})
}
