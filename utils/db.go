package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
)

var supabaseClient *supabase.Client

func InitSupabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")

	client, err := supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{
		Schema: "public",
	})
	if client == nil {
		log.Fatal("Erro ao inicializar o cliente Supabase", err)
	}

	supabaseClient = client
}

func GetSupabase() *supabase.Client {
	return supabaseClient
}
