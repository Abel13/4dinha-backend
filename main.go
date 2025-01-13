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
	supabaseServiceRole := os.Getenv("SUPABASE_SERVICE_ROLE")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}

	if jwtSecret == "" || supabaseURL == "" || supabaseKey == "" || supabaseServiceRole == "" {
		log.Fatal("Configuração de ambiente inválida: verifique JWT_SECRET, SUPABASE_URL e SUPABASE_ANON_KEY")
	}

	authService := services.NewAuthService(jwtSecret)
	middlewareAuth := middleware.NewAuthMiddleware(authService, supabaseURL, supabaseKey)

	matchRepo := repositories.NewMatchRepository()
	serviceMatchRepo := repositories.NewServiceMatchRepository(supabaseServiceRole, supabaseURL, supabaseKey)
	matchUserRepo := repositories.NewMatchUsersRepository()
	serviceMatchUsersRepo := repositories.NewServiceMatchUsersRepository(supabaseServiceRole, supabaseURL, supabaseKey)
	matchActionRepo := repositories.NewMatchActionRepository()
	deckRepo := repositories.NewDeckRepository()
	playerCardRepo := repositories.NewPlayerCardsRepository()
	servicePlayerCardsRepo := repositories.NewServicePlayerCardsRepository(supabaseServiceRole, supabaseURL, supabaseKey)
	roundRepo := repositories.NewRoundRepo()
	betRepo := repositories.NewBetRepository()

	trumpService := services.NewTrumpService(*matchRepo, *roundRepo, *deckRepo)
	dealService := services.NewDealService(*matchRepo, *matchUserRepo, *matchActionRepo, *deckRepo, *playerCardRepo, *roundRepo)
	updateService := services.NewUpdateService(
		*matchRepo,
		*matchUserRepo,
		*servicePlayerCardsRepo,
		*roundRepo,
		*deckRepo,
		*matchActionRepo,
		*betRepo)
	playService := services.NewPlayService(*playerCardRepo, *roundRepo)
	roundService := services.NewRoundService(*serviceMatchRepo, *roundRepo, *servicePlayerCardsRepo, *deckRepo, *serviceMatchUsersRepo, *betRepo)

	dealHandler := handlers.NewDealHandler(dealService)
	updateHandler := handlers.NewUpdateHandler(updateService)
	playHandler := handlers.NewPlayHandler(playService)
	trumpHandler := handlers.NewTrumpHandler(trumpService)
	roundHandler := handlers.NewRoundHandler(roundService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	protected := router.Group("/api")
	protected.Use(middlewareAuth.HandleAuth)
	{
		protected.POST("/deal", dealHandler.DealCards)
		protected.POST("/play", playHandler.Play)
		protected.GET("/update", updateHandler.Update)
		protected.GET("/trumps", trumpHandler.Trumps)
		protected.PUT("/finish-round", roundHandler.FinishRound)
	}

	log.Printf("Servidor rodando na porta %s", port)
	router.Run(":" + port)
}
