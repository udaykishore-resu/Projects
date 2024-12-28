package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mock points data - in a real application, this would be in a database
var userPoints = map[string]int{
	"1": 100,
}

func GetUserPoints(c *gin.Context) {
	userId := c.Param("userId")
	points, exists := userPoints[userId]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points})
}

func AddPoints(c *gin.Context) {
	var req struct {
		UserID string `json:"userId" binding:"required"`
		Points int    `json:"points" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentPoints := userPoints[req.UserID]
	userPoints[req.UserID] = currentPoints + req.Points

	c.JSON(http.StatusOK, gin.H{"points": userPoints[req.UserID]})
}
