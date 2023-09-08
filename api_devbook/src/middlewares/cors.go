package middlewares

import (
	"fmt"
	"net/http"
)

// This middleware is not being used because the CORS is being handled by the package "github.com/rs/cors"
// This package is being used in the src/server/server.go file
func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CORS Middleware")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Controll-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
		next(w, r)
	}
}
