package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.AuthTokenValidate(r); err != nil {
			responses.ResponseError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
