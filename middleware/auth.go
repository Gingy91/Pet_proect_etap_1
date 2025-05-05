package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var jwtKey = []byte("secret")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missed", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Не правильный токен", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
