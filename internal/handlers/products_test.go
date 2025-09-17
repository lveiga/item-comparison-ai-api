package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"item-comparison-ai-api/internal/database"
	"item-comparison-ai-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockDB is a mock implementation of the DB interface
type MockDB struct {
	products []models.Product
}

// LoadProducts returns the products from the mock database
func (m *MockDB) LoadProducts() ([]models.Product, error) {
	return m.products, nil
}

// SaveProducts saves the products to the mock database
func (m *MockDB) SaveProducts(products []models.Product) error {
	m.products = products
	return nil
}

// GetNextID calculates the next available ID for a new product
func (m *MockDB) GetNextID(products []models.Product) int {
	maxID := 0
	for _, p := range products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	return maxID + 1
}

func setupTestRouter(db database.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewProductHandler(db)
	r.GET("/products", h.GetAllProducts)
	r.GET("/products/:id", h.GetProduct)
	r.POST("/products", h.CreateProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.PATCH("/products/:id", h.PatchProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	return r
}

func TestGetProduct(t *testing.T) {
	db := &MockDB{
		products: []models.Product{
			{ID: 1, Name: "Laptop", Category: "Electronics"},
		},
	}
	r := setupTestRouter(db)

	// Test case 1: Success scenario
	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

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
		assert.JSONEq(t, `{"error":"Product not found"}`, w.Body.String())
	})

	// Test case 3: Invalid ID
	t.Run("InvalidID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/products/abc", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error":"Invalid product ID"}`, w.Body.String())
	})
}

func TestGetAllProducts(t *testing.T) {
	db := &MockDB{
		products: []models.Product{
			{ID: 1, Name: "Laptop"},
			{ID: 2, Name: "Smartphone"},
		},
	}
	r := setupTestRouter(db)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var prods []models.Product
	json.Unmarshal(w.Body.Bytes(), &prods)
	assert.Len(t, prods, 2)
}

func TestCreateProduct(t *testing.T) {
	db := &MockDB{
		products: []models.Product{},
	}
	r := setupTestRouter(db)

	newProduct := models.Product{
		Name:     "New Product",
		Category: "Electronics",
	}
	jsonProduct, _ := json.Marshal(newProduct)

	req, _ := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonProduct))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdProduct models.Product
	json.Unmarshal(w.Body.Bytes(), &createdProduct)

	assert.Equal(t, 1, createdProduct.ID)
	assert.Equal(t, "New Product", createdProduct.Name)

	products, _ := db.LoadProducts()
	assert.Len(t, products, 1)
}

func TestUpdateProduct(t *testing.T) {
	db := &MockDB{
		products: []models.Product{
			{ID: 1, Name: "Laptop", Category: "Electronics"},
		},
	}
	r := setupTestRouter(db)

	updatedProduct := models.Product{
		Name:     "Updated Laptop",
		Category: "Premium Electronics",
	}
	jsonProduct, _ := json.Marshal(updatedProduct)

	req, _ := http.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(jsonProduct))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var returnedProduct models.Product
	json.Unmarshal(w.Body.Bytes(), &returnedProduct)
	assert.Equal(t, "Updated Laptop", returnedProduct.Name)
	assert.Equal(t, "Premium Electronics", returnedProduct.Category)

	products, _ := db.LoadProducts()
	assert.Equal(t, "Updated Laptop", products[0].Name)
}

func TestPatchProduct(t *testing.T) {
	db := &MockDB{
		products: []models.Product{
			{ID: 1, Name: "Laptop", Category: "Electronics"},
		},
	}
	r := setupTestRouter(db)

	patchData := map[string]interface{}{
		"category": "Premium Electronics",
	}
	jsonPatch, _ := json.Marshal(patchData)

	req, _ := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewBuffer(jsonPatch))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var returnedProduct models.Product
	json.Unmarshal(w.Body.Bytes(), &returnedProduct)
	assert.Equal(t, "Premium Electronics", returnedProduct.Category)

	products, _ := db.LoadProducts()
	assert.Equal(t, "Premium Electronics", products[0].Category)
}

func TestDeleteProduct(t *testing.T) {
	db := &MockDB{
		products: []models.Product{
			{ID: 1, Name: "Laptop"},
		},
	}
	r := setupTestRouter(db)

	req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	products, _ := db.LoadProducts()
	assert.Len(t, products, 0)
}
