package middlewares

import (
	"net/http"
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

// func Authenticate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			if r.Header["Token"] == nil {
// 				var err error
// 				response.Error(w, http.StatusUnauthorized, err)
// 			}
// 		})
// }
