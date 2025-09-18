package routes

import (
	"item-comparison-ai-api/config"
	"item-comparison-ai-api/internal/database"
	"item-comparison-ai-api/internal/handlers"
	"item-comparison-ai-api/internal/repositories"
	"item-comparison-ai-api/internal/server"

	"github.com/gin-gonic/gin"
)

// Handler - represents a route/controller binder
type ProductRouter struct{}

// Bind - method responsible to bind controller and actions
func (r *ProductRouter) Bind(router *gin.RouterGroup, app *server.Application) {
	db := database.NewClient(&database.Database{})

	baseRepo := repositories.NewBaseRepository(db, config.New())
	productRepo := repositories.NewProductRepository(baseRepo)
	productHandler := handlers.NewProductHandler(productRepo)

	// Define the GET endpoint for retrieving a product by ID
	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/:id", productHandler.GetProduct)
	router.POST("/products", productHandler.CreateProduct)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.PATCH("/products/:id", productHandler.PatchProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)
}
