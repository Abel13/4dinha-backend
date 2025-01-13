package main

import (
	"4dinha-backend/handlers"
	"4dinha-backend/middleware"
	"4dinha-backend/repositories"
	"4dinha-backend/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	if jwtSecret == "" || supabaseURL == "" || supabaseKey == "" {
		log.Fatal("Configuração de ambiente inválida: verifique JWT_SECRET, SUPABASE_URL e SUPABASE_ANON_KEY")
	}

	authService := services.NewAuthService(jwtSecret)
	middlewareAuth := middleware.NewAuthMiddleware(authService, supabaseURL, supabaseKey)

	matchRepo := repositories.NewMatchRepository()
	matchUserRepo := repositories.NewMatchUsersRepository()
	deckRepo := repositories.NewDeckRepository()

	dealService := services.NewDealService(*matchRepo, *matchUserRepo, *deckRepo)

	dealHandler := handlers.NewDealHandler(dealService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	protected := router.Group("/api")
	protected.Use(middlewareAuth.HandleAuth)
	{
		protected.POST("/deal", dealHandler.DealCards)
	}

	log.Printf("Servidor rodando na porta %s", port)
	router.Run(":" + port)
}
