package handlers

import (
	"4dinha-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
	"net/http"
)

type TrumpHandler struct {
	TrumpService *services.TrumpService
}

func NewTrumpHandler(trumpService *services.TrumpService) *TrumpHandler {
	return &TrumpHandler{
		TrumpService: trumpService,
	}
}

func (h *TrumpHandler) Trumps(c *gin.Context) {
	matchID := c.Query("matchID")
	if matchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Match ID not provided"})
		return
	}

	client, exists := c.Get("supabaseClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Supabase client not found in context"})
		return
	}

	supabaseClient, ok := client.(*supabase.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Supabase client"})
		return
	}

	trumps, err := h.TrumpService.GetTrumps(supabaseClient, matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trumps)
}
