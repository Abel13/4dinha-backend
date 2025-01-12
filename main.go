package main

import (
	"4dinha-backend/handlers"
	"4dinha-backend/middleware"
	"4dinha-backend/repositories"
	"4dinha-backend/services"
	"4dinha-backend/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	// Inicializa o Supabase
	utils.InitSupabase()
	db := utils.GetSupabase()

	// Configurações
	supabaseAnonKey := os.Getenv("SUPABASE_ANON_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333" // Valor padrão
	}

	// Inicializa os repositórios
	matchRepo := &repositories.SupabaseMatchRepository{DB: db}
	matchUserRepo := &repositories.SupabaseMatchUsersRepository{DB: db}
	deckRepo := &repositories.SupabaseDeckRepository{DB: db}

	// Inicializa os serviços
	authService := services.NewAuthService(supabaseAnonKey)
	dealService := services.NewDealService(matchRepo, matchUserRepo, deckRepo)

	// Inicializa os handlers
	dealHandler := handlers.NewDealHandler(dealService)

	// Middleware de autenticação
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Configura o roteador
	router := gin.Default()

	// Rotas públicas
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is running!"})
	})

	// Rotas protegidas
	protected := router.Group("/api")
	protected.Use(authMiddleware.HandleAuth) // Aplica o middleware de autenticação
	{
		protected.POST("/deal", dealHandler.DealCards)
	}

	// Inicia o servidor
	router.Run(":" + port)
}
