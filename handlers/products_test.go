package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"item-comparison-ai-api/models"
)

func TestGetProduct(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup the router
	r := gin.Default()
	r.GET("/products/:id", GetProduct)

	// Test case 1: Success scenario
	t.Run("Success", func(t *testing.T) {
		// Create a request to pass to our handler
		req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Check the status code is what we expect
		assert.Equal(t, http.StatusOK, w.Code)

		// Check the response body is what we expect
		var product models.Product
		json.Unmarshal(w.Body.Bytes(), &product)

		assert.Equal(t, 1, product.ID)
		assert.Equal(t, "Laptop", product.Name)
	})

	// Test case 2: Product not found
	t.Run("NotFound", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/products/999", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Product not found", response["error"])
	})

	// Test case 3: Invalid ID format
	t.Run("InvalidID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/products/abc", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "Invalid product ID", response["error"])
	})
}
