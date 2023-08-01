package middleware

import (
	"fmt"
	"net/http"
)

var apiKey = "1234"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get("X-API-Key")
		if apiKeyHeader == "" || apiKeyHeader != apiKey {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized: Invalid API key")
			return
		}

		next.ServeHTTP(w, r)
	})
}
