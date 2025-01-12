package handlers

import (
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
	var body struct {
		MatchID string `json:"matchId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.DealService.DealCards(c.GetHeader("Authorization"), body.MatchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cards dealt successfully"})
}
