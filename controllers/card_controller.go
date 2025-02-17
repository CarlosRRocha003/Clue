package controllers

import (
	"Clue/models/card"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func GetAllCards(c *gin.Context) {
	cards := card.GetAllCards()
	c.JSON(http.StatusOK, cards)
}

func GetCategories(c *gin.Context) {
	categories := card.GetCategories()
	c.JSON(http.StatusOK, categories)
}

func AddCard(c *gin.Context) {
	var newCard card.Card
	if err := c.ShouldBindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card.AddCard(newCard)

	c.JSON(http.StatusCreated, gin.H{"message": "Card added successfully"})
}
