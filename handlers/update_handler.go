package handlers

import (
	"4dinha-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
	"net/http"
)

type UpdateHandler struct {
	UpdateService *services.UpdateService
}

func NewUpdateHandler(updateService *services.UpdateService) *UpdateHandler {
	return &UpdateHandler{UpdateService: updateService}
}

func (h *UpdateHandler) Update(c *gin.Context) {
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

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	game, err := h.UpdateService.Update(supabaseClient, matchID, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}
