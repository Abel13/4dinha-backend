package middleware

import (
	"4dinha-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	AuthService *services.AuthService
}

func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{AuthService: authService}
}

func (m *AuthMiddleware) HandleAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	userId, err := m.AuthService.ValidateToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("userID", userId)
	c.Next()
}
