package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

type AuthService struct {
	SupabaseAnonKey string
}

func NewAuthService(supabaseAnonKey string) *AuthService {
	return &AuthService{SupabaseAnonKey: supabaseAnonKey}
}

func (s *AuthService) ValidateToken(authHeader string) (string, error) {
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	tokenString := tokenParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(s.SupabaseAnonKey), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
		return "", errors.New("missing sub claim in token")
	}

	return "", errors.New("invalid token claims")
}
