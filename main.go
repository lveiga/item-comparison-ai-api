package main

import (
	"github.com/gin-gonic/gin"
	"item-comparison-ai-api/handlers"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Define the GET endpoint for retrieving a product by ID
	r.GET("/products/:id", handlers.GetProduct)

	// Run the server on port 8080
	r.Run(":8080")
}
