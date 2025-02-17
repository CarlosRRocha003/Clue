package controllers

import "github.com/gin-gonic/gin"

// RegisterRoutes sets up the routes for the handlers package.
func RegisterRoutes(r *gin.Engine) {
	RegisterCardRoutes(r)
}

func RegisterCardRoutes(r *gin.Engine) {
	r.GET("/hello", HelloHandler)
	r.GET("/allCards", GetAllCards)
	r.GET("/categories", GetCategories)
	r.GET("/allPlayers", GetAllPlayers)
	r.POST("/addCard", AddCard)
	r.POST("/addPlayer", AddPlayer)
}
