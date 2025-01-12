package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var secret = "XkToCmc+QjBIrpnB1WldYBuXMYsEwUBiiMDOXsxrUqUk04gGV+NEeSNaKZVo8Q++28qF9haseyH+xpXRzW2+dg=="

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"].(string) == "" {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Adiciona o user_id ao contexto
		ctx := context.WithValue(r.Context(), "user_id", claims["sub"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
