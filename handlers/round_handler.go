package handlers

import (
	"4dinha-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
	"net/http"
)

type RoundHandler struct {
	RoundService *services.RoundService
}

func NewRoundHandler(roundService *services.RoundService) *RoundHandler {
	return &RoundHandler{RoundService: roundService}
}

func (h *RoundHandler) FinishRound(c *gin.Context) {
	var body struct {
		MatchID string `json:"matchId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	client, exists := c.Get("supabaseClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Supabase client not found in context"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	supabaseClient, ok := client.(*supabase.Client)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Supabase client"})
		return
	}

	err := h.RoundService.FinishRound(supabaseClient, body.MatchID, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success"})
}
