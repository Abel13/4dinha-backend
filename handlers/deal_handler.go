package handlers

import (
	"github.com/supabase-community/supabase-go"
	"net/http"

	"4dinha-backend/services"

	"github.com/gin-gonic/gin"
)

type DealHandler struct {
	DealService *services.DealService
}

func NewDealHandler(dealService *services.DealService) *DealHandler {
	return &DealHandler{DealService: dealService}
}

func (h *DealHandler) DealCards(c *gin.Context) {
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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	var body struct {
		MatchID string `json:"matchId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.DealService.DealCards(supabaseClient, userID.(string), body.MatchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cards dealt successfully"})
}
