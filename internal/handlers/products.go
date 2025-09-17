package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"item-comparison-ai-api/internal/models"
)

// products slice to seed product data.
var products = []models.Product{
	{ID: 1, Name: "Laptop", ImageURL: "/images/laptop.png", Description: "High-performance laptop", Price: 1200.00, Rating: 4.5, Specifications: map[string]string{"RAM": "16GB", "Storage": "512GB SSD"}},
	{ID: 2, Name: "Smartphone", ImageURL: "/images/smartphone.png", Description: "Latest model smartphone", Price: 800.00, Rating: 4.8, Specifications: map[string]string{"Camera": "108MP", "Battery": "5000mAh"}},
	{ID: 3, Name: "Headphones", ImageURL: "/images/headphones.png", Description: "Noise-cancelling headphones", Price: 150.00, Rating: 4.2, Specifications: map[string]string{"Connectivity": "Bluetooth 5.0", "Driver size": "40mm"}},
}

// GetProduct retrieves a product by its ID
func GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
