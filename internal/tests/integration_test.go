package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"item-comparison-ai-api/internal/handlers"
	"item-comparison-ai-api/internal/models"
)

// setupRouter initializes the Gin router for testing
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/products/:id", handlers.GetProduct)
	return r
}

// TestIntegrationGetProductSuccess tests the success scenario for the /products/{id} endpoint
func TestIntegrationGetProductSuccess(t *testing.T) {
	t.Parallel()

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
	body, err := ioutil.ReadAll(resp.Body)
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
	t.Parallel()

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
