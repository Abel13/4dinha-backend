package middleware

import (
	"4dinha-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
)

type AuthMiddleware struct {
	authService *services.AuthService
	supabaseURL string
	supabaseKey string
}

func NewAuthMiddleware(authService *services.AuthService, supabaseURL, supabaseKey string) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		supabaseURL: supabaseURL,
		supabaseKey: supabaseKey,
	}
}

func (m *AuthMiddleware) HandleAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	user, err := m.authService.ValidateToken(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	client, err := supabase.NewClient(m.supabaseURL, m.supabaseKey, &supabase.ClientOptions{
		Headers: map[string]string{
			"Authorization": authHeader,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Supabase client"})
		c.Abort()
		return
	}

	c.Set("supabaseClient", client)
	c.Set("userID", user)

	c.Next()
}
