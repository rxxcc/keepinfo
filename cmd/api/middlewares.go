package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/inuoshios/keepinfo/internal/auth"
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

var (
	authTypeBearer = "Bearer"
	authPayloadKey = "AuthorizationPayload"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

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

			payload, err := auth.VerifyToken(accessToken)
			if err != nil {
				response.Error(w, http.StatusUnauthorized, fmt.Errorf("-> %w", err))
				return
			}

			ctx := context.WithValue(r.Context(), authPayloadKey, payload)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
