package controllers

import (
	"Clue/models/player"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPlayers(c *gin.Context) {
	players := player.GetAllPlayers()
	c.JSON(http.StatusOK, players)
}

func AddPlayer(c *gin.Context) {
	var newPlayer player.Player

	if err := c.ShouldBindJSON(&newPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player.AddPlayer(newPlayer)

	c.JSON(http.StatusCreated, gin.H{"message": "Player added successfully"})
}
