package services

import (
	"4dinha-backend/models"
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

func (s *AuthService) ValidateToken(authHeader string) (*models.User, error) {
	// Verifica se o cabeçalho está presente
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header format")
	}

	tokenString := tokenParts[1]

	// Analisa e valida o token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Retorna a chave secreta
		return []byte(s.SupabaseAnonKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extrai as claims e retorna os dados do usuário
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user := &models.User{
			ID:    claims["sub"].(string),
			Email: claims["email"].(string),
		}
		return user, nil
	}

	return nil, errors.New("invalid token claims")
}
