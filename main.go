package main

import (
	"Clue/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Register routes using the views package
	controllers.RegisterRoutes(r)

	// Serve other static files from "/static"
	r.Static("/static", "./view")

	// Start the server on port 8080
	fmt.Println("Server running on http://localhost:8080/")
	r.Run(":8080")
}
