package repositories

import (
	"encoding/json"
	"item-comparison-ai-api/internal/models"
	"os"
)

type ProductRepository interface {
	LoadProducts() ([]models.Product, error)
	SaveProducts([]models.Product) error
	GetNextID([]models.Product) int
}

// productRepository implements the ProductRepository interface
type productRepository struct {
	baseRepo BaseRepositoryInterface
}

// LoadProducts reads products from the data.json file
func (p *productRepository) LoadProducts() ([]models.Product, error) {
	data, err := p.baseRepo.Load()
	if err != nil {
		if os.IsNotExist(err) {
			return make([]models.Product, 0), nil // Return empty slice if file doesn't exist
		}
		return nil, err
	}

	var products []models.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	// Convert []models.Product to []models.Product
	result := make([]models.Product, len(products))
	for i, p := range products {
		result[i] = p
	}

	return result, nil
}

// SaveProducts writes products to the data.json file
func (p *productRepository) SaveProducts(model []models.Product) error {
	data := make([]interface{}, len(model))
	for idx, item := range model {
		data[idx] = item
	}

	return p.baseRepo.Save(data)
}

// GetNextID calculates the next available ID for a new product
func (p *productRepository) GetNextID(products []models.Product) int {
	maxID := 0
	for _, pr := range products {
		if pr.ID > maxID {
			maxID = pr.ID
		}
	}
	return maxID + 1
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(baseRepo BaseRepositoryInterface) ProductRepository {
	return &productRepository{baseRepo: baseRepo}
}
