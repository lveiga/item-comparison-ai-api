package main

import (
	"github.com/gin-gonic/gin"
	"item-comparison-ai-api/internal/database"
	"item-comparison-ai-api/internal/handlers"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Create a new database client
	db := &database.Client{}

	// Create a new product handler
	productHandler := handlers.NewProductHandler(db)

	// Define the GET endpoint for retrieving a product by ID
	r.GET("/products", productHandler.GetAllProducts)
	r.GET("/products/:id", productHandler.GetProduct)
	r.POST("/products", productHandler.CreateProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.PATCH("/products/:id", productHandler.PatchProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	// Run the server on port 8080
	r.Run(":8080")
}
