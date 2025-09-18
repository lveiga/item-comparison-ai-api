package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"item-comparison-ai-api/config"
	"item-comparison-ai-api/internal/database"
	"item-comparison-ai-api/internal/handlers"
	"item-comparison-ai-api/internal/models"
	"item-comparison-ai-api/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestEnvironment creates a temporary data.json for testing
func setupTestEnvironment(t *testing.T) func() {
	// Create a temporary data.json file
	tempFile, err := os.CreateTemp("", "test_data_*.json")
	assert.NoError(t, err)
	tempFileName := tempFile.Name()
	tempFile.Close()

	t.Setenv("DATA_FILE_PATH", tempFileName)

	// Seed initial data
	initialProducts := []models.Product{
		{ID: 1, Name: "Laptop", ImageURL: "/images/laptop.png", Description: "High-performance laptop", Price: 1200.00, Rating: 4.5, Specifications: map[string]string{"RAM": "16GB", "Storage": "512GB SSD"}, Category: "Electronics"},
		{ID: 2, Name: "Smartphone", ImageURL: "/images/smartphone.png", Description: "Latest model smartphone", Price: 800.00, Rating: 4.8, Specifications: map[string]string{"Camera": "108MP", "Battery": "5000mAh"}, Category: "Electronics"},
		{ID: 3, Name: "Headphones", ImageURL: "/images/headphones.png", Description: "Noise-cancelling headphones", Price: 150.00, Rating: 4.2, Specifications: map[string]string{"Connectivity": "Bluetooth 5.0", "Driver size": "40mm"}, Category: "Accessories"},
	}
	var config = config.New()

	db := database.NewClient(&database.Database{})
	baseRepo := repositories.NewBaseRepository(db, config)
	repo := repositories.NewProductRepository(baseRepo)
	err = repo.SaveProducts(initialProducts)
	assert.NoError(t, err)

	// Return a cleanup function
	return func() {
		os.Remove(tempFileName) // Clean up the temporary file
	}
}

// setupRouter initializes the Gin router for testing
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	db := database.NewClient(&database.Database{})
	baseRepo := repositories.NewBaseRepository(db, config.New())
	repo := repositories.NewProductRepository(baseRepo)
	h := handlers.NewProductHandler(repo)
	r.GET("/products", h.GetAllProducts)
	r.GET("/products/:id", h.GetProduct)
	r.POST("/products", h.CreateProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.PATCH("/products/:id", h.PatchProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	return r
}

// TestIntegrationGetProductSuccess tests the success scenario for the /products/{id} endpoint
func TestIntegrationGetProductSuccess(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	// Setup the router
	router := setupRouter()

	// Create a new test server
	server := httptest.NewServer(router)
	defer server.Close()

	// Make a GET request to the test server
	resp, err := http.Get(server.URL + "/products/1")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert the status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Unmarshal the response body into a Product struct
	var product models.Product
	err = json.Unmarshal(body, &product)
	assert.NoError(t, err)

	// Assert the product details
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Laptop", product.Name)
}

// TestIntegrationNotFound tests the scenario where the endpoint is not found
func TestIntegrationNotFound(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	// Setup the router
	router := setupRouter()

	// Create a new test server
	server := httptest.NewServer(router)
	defer server.Close()

	// Make a GET request to a non-existent endpoint
	resp, err := http.Get(server.URL + "/api/v1/products/1")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert the status code
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

// TestIntegrationGetAllProducts tests the GetAllProducts endpoint
func TestIntegrationGetAllProducts(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/products?limit=1&offset=1")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var products []models.Product
	err = json.Unmarshal(body, &products)
	assert.NoError(t, err)

	assert.Len(t, products, 1)
	assert.Equal(t, 2, products[0].ID)
}

// TestIntegrationCreateProduct tests the CreateProduct endpoint
func TestIntegrationCreateProduct(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	newProduct := models.Product{
		Name:        "New Product",
		Description: "A brand new product",
		Price:       100.00,
		Rating:      4.0,
		Category:    "Electronics",
	}
	jsonProduct, _ := json.Marshal(newProduct)

	resp, err := http.Post(server.URL+"/products", "application/json", bytes.NewBuffer(jsonProduct))
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var createdProduct models.Product
	err = json.Unmarshal(body, &createdProduct)
	assert.NoError(t, err)

	assert.Equal(t, "New Product", createdProduct.Name)
	assert.Equal(t, 4, createdProduct.ID)
}

// TestIntegrationUpdateProduct tests the UpdateProduct endpoint
func TestIntegrationUpdateProduct(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	updatedProduct := models.Product{
		ID:          1,
		Name:        "Updated Laptop",
		Description: "Updated description",
		Price:       1300.00,
		Rating:      4.6,
		Category:    "Electronics",
	}
	jsonProduct, _ := json.Marshal(updatedProduct)

	req, _ := http.NewRequest(http.MethodPut, server.URL+"/products/1", bytes.NewBuffer(jsonProduct))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var returnedProduct models.Product
	err = json.Unmarshal(body, &returnedProduct)
	assert.NoError(t, err)

	assert.Equal(t, "Updated Laptop", returnedProduct.Name)
}

// TestIntegrationPatchProduct tests the PatchProduct endpoint
func TestIntegrationPatchProduct(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	patchData := map[string]interface{}{
		"price":    1250.00,
		"rating":   4.7,
		"category": "Premium Electronics",
	}
	jsonPatch, _ := json.Marshal(patchData)

	req, _ := http.NewRequest(http.MethodPatch, server.URL+"/products/1", bytes.NewBuffer(jsonPatch))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var returnedProduct models.Product
	err = json.Unmarshal(body, &returnedProduct)
	assert.NoError(t, err)

	assert.Equal(t, 1250.00, returnedProduct.Price)
	assert.Equal(t, 4.7, returnedProduct.Rating)
	assert.Equal(t, "Premium Electronics", returnedProduct.Category)
}

// TestIntegrationDeleteProduct tests the DeleteProduct endpoint
func TestIntegrationDeleteProduct(t *testing.T) {
	cleanup := setupTestEnvironment(t)
	defer cleanup()

	router := setupRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	req, _ := http.NewRequest(http.MethodDelete, server.URL+"/products/1", nil)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
