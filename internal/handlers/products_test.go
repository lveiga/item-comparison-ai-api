package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"item-comparison-ai-api/internal/models"
	"item-comparison-ai-api/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of the ProductRepository interface
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) LoadProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) SaveProducts(products []models.Product) error {
	args := m.Called(products)
	return args.Error(0)
}

func (m *MockProductRepository) GetNextID(products []models.Product) int {
	args := m.Called(products)
	return args.Int(0)
}

func setupTestRouter(repo repositories.ProductRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewProductHandler(repo)
	r.GET("/products", h.GetAllProducts)
	r.GET("/products/:id", h.GetProduct)
	r.POST("/products", h.CreateProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.PATCH("/products/:id", h.PatchProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	return r
}

func TestGetProduct(t *testing.T) {
	// Test case 1: Success scenario
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("LoadProducts").Return([]models.Product{{ID: 1, Name: "Laptop", Category: "Electronics"}}, nil)

		r := setupTestRouter(mockRepo)
		req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var product models.Product
		json.Unmarshal(w.Body.Bytes(), &product)

		assert.Equal(t, 1, product.ID)
		assert.Equal(t, "Laptop", product.Name)
		mockRepo.AssertExpectations(t)
	})

	// Test case 2: Product not found
	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("LoadProducts").Return([]models.Product{{ID: 1, Name: "Laptop"}}, nil)

		r := setupTestRouter(mockRepo)
		req, _ := http.NewRequest(http.MethodGet, "/products/999", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"error":"Not found"}`, w.Body.String())
		mockRepo.AssertExpectations(t)
	})

	// Test case 3: Invalid ID
	t.Run("InvalidID", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		r := setupTestRouter(mockRepo)
		req, _ := http.NewRequest(http.MethodGet, "/products/abc", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error":"Invalid ID"}`, w.Body.String())
	})
}

func TestGetAllProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockRepo.On("LoadProducts").Return([]models.Product{{ID: 1, Name: "Laptop"}, {ID: 2, Name: "Smartphone"}}, nil)

	r := setupTestRouter(mockRepo)
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var prods []models.Product
	json.Unmarshal(w.Body.Bytes(), &prods)
	assert.Len(t, prods, 2)
	mockRepo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	newProduct := models.Product{Name: "New Product", Category: "Electronics"}

	mockRepo.On("LoadProducts").Return([]models.Product{}, nil)
	mockRepo.On("GetNextID", mock.Anything).Return(1)
	mockRepo.On("SaveProducts", mock.Anything).Return(nil)

	r := setupTestRouter(mockRepo)

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
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	updatedProduct := models.Product{Name: "Updated Laptop", Category: "Premium Electronics"}

	mockRepo.On("LoadProducts").Return([]models.Product{{ID: 1, Name: "Laptop", Category: "Electronics"}}, nil)
	mockRepo.On("SaveProducts", mock.Anything).Return(nil)

	r := setupTestRouter(mockRepo)

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
	mockRepo.AssertExpectations(t)
}

func TestPatchProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	p := models.Product{ID: 1, Name: "Laptop", Category: "Electronics"}

	mockRepo.On("LoadProducts").Return([]models.Product{p}, nil)
	p.Category = "Premium Electronics"
	mockRepo.On("SaveProducts", mock.Anything).Return(nil)

	r := setupTestRouter(mockRepo)

	patchData := map[string]interface{}{"category": "Premium Electronics"}
	jsonPatch, _ := json.Marshal(patchData)

	req, _ := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewBuffer(jsonPatch))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var returnedProduct models.Product
	json.Unmarshal(w.Body.Bytes(), &returnedProduct)
	assert.Equal(t, "Premium Electronics", returnedProduct.Category)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockRepo.On("LoadProducts").Return([]models.Product{{ID: 1, Name: "Laptop"}}, nil)
	mockRepo.On("SaveProducts", mock.Anything).Return(nil)

	r := setupTestRouter(mockRepo)

	req, _ := http.NewRequest(http.MethodDelete, "/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}
