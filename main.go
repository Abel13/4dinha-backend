package main

import (
	"4dinha-backend/handlers"
	"4dinha-backend/middleware"
	"4dinha-backend/repositories"
	"4dinha-backend/services"
	"4dinha-backend/utils"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}
	utils.InitSupabase()
	db := utils.GetSupabase()

	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	matchRepo := &repositories.SupabaseMatchRepository{DB: db}
	matchUserRepo := &repositories.SupabaseMatchUsersRepository{DB: db}
	deckRepo := &repositories.SupabaseDeckRepository{DB: db}

	authService := services.NewAuthService(jwtSecret)
	dealService := services.NewDealService(matchRepo, matchUserRepo, deckRepo)

	dealHandler := handlers.NewDealHandler(dealService)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	protected := router.Group("/api")
	protected.Use(authMiddleware.HandleAuth)
	{
		protected.POST("/deal", dealHandler.DealCards)
	}

	router.Run(":" + port)
}
