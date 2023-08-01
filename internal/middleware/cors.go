package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow GET, POST, PUT, DELETE methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		// Allow Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
