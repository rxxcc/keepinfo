package middlewares

import (
	"net/http"

	"github.com/0xmlx/keepinfo/internal/config"
)

var sessionManager config.Config

// AddContentType functions helps set headers for our api.
func AddContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		next.ServeHTTP(w, r)
	})
}

// LoadSession automatically load and saves the session data for the current request.
func LoadSession(next http.Handler) http.Handler {
	return sessionManager.Session.LoadAndSave(next)
}
