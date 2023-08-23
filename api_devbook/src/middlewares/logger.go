package middlewares

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s | %s%s]", r.Method, r.Host, r.RequestURI)
		next(w, r)
	}
}
