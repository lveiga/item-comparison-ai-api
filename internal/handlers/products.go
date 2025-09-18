package handlers

import (
	"net/http"
	"strconv"

	"item-comparison-ai-api/internal/models"
	"item-comparison-ai-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

// ProductHandler holds the database client
type ProductHandler struct {
	repo repositories.ProductRepository
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(repository repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repository}
}

// GetProduct retrieves a product by its ID
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, ErrInvalidID)
		return
	}

	products, err := h.repo.LoadProducts()
	if err != nil {
		HandleError(c, ErrFailedToLoad)
		return
	}

	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	HandleError(c, ErrNotFound)
}

// GetAllProducts retrieves all products with optional pagination
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.repo.LoadProducts()
	if err != nil {
		HandleError(c, ErrFailedToLoad)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		HandleError(c, ErrInvalidLimitParameter)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		HandleError(c, ErrInvalidOffsetParameter)
		return
	}

	start := offset
	end := offset + limit

	if start > len(products) {
		start = len(products)
	}
	if end > len(products) {
		end = len(products)
	}

	c.JSON(http.StatusOK, products[start:end])
}

// CreateProduct adds a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		HandleError(c, ErrBindJSON)
		return
	}

	products, err := h.repo.LoadProducts()
	if err != nil {
		HandleError(c, ErrFailedToLoad)
		return
	}

	newProduct.ID = h.repo.GetNextID(products)
	products = append(products, newProduct)

	if err := h.repo.SaveProducts(products); err != nil {
		HandleError(c, ErrFailedToSave)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

// UpdateProduct updates an existing product by ID
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, ErrInvalidID)
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		HandleError(c, ErrBindJSON)
		return
	}

	products, err := h.repo.LoadProducts()
	if err != nil {
		HandleError(c, ErrFailedToLoad)
		return
	}

	found := false
	for i, p := range products {
		if p.ID == id {
			updatedProduct.ID = id // Ensure the ID from the URL is used
			products[i] = updatedProduct
			found = true
			break
		}
	}

	if !found {
		HandleError(c, ErrNotFound)
		return
	}

	if err := h.repo.SaveProducts(products); err != nil {
		HandleError(c, ErrFailedToSave)
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// PatchProduct partially updates an existing product by ID
func (h *ProductHandler) PatchProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, ErrInvalidID)
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		HandleError(c, ErrBindJSON)
		return
	}

	products, err := h.repo.LoadProducts()
	if err != nil {
		HandleError(c, ErrFailedToLoad)
		return
	}

	found := false
	for i, p := range products {
		if p.ID == id {
			if name, ok := updates["name"]; ok {
				p.Name = name.(string)
			}
			if imageURL, ok := updates["image_url"]; ok {
				p.ImageURL = imageURL.(string)
			}
			if description, ok := updates["description"]; ok {
				p.Description = description.(string)
			}
			if price, ok := updates["price"]; ok {
				p.Price = price.(float64)
			}
			if rating, ok := updates["rating"]; ok {
				p.Rating = rating.(float64)
			}
			if specs, ok := updates["specifications"]; ok {
				if specMap, isMap := specs.(map[string]interface{}); isMap {
					convertedSpecs := make(map[string]string)
					for k, v := range specMap {
						if strVal, isString := v.(string); isString {
							convertedSpecs[k] = strVal
						}
					}
					p.Specifications = convertedSpecs
				}
			}
			if category, ok := updates["category"]; ok {
				p.Category = category.(string)
			}
			products[i] = p
			found = true
			break
		}
	}

	if !found {
		HandleError(c, ErrNotFound)
		return
	}

	if err := h.repo.SaveProducts(products); err != nil {
		HandleError(c, ErrFailedToSave)
		return
	}

	c.JSON(http.StatusOK, products[id-1]) // Assuming IDs are sequential for simplicity
}

// DeleteProduct removes a product by ID
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, ErrInvalidID)
		return
	}

	products, err := h.repo.LoadProducts()
	if err != nil {
		HandleError(c, ErrFailedToLoad)
		return
	}

	found := false
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		HandleError(c, ErrNotFound)
		return
	}

	if err := h.repo.SaveProducts(products); err != nil {
		HandleError(c, ErrFailedToSave)
		return
	}

	c.Status(http.StatusNoContent)
}
