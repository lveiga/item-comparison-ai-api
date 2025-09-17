package database

import (
	"encoding/json"
	"os"
	"sync"

	"item-comparison-ai-api/internal/models"
)

var DataFile = os.Getenv("DATA_FILE_PATH")

var (
	mu sync.Mutex // Mutex to protect access to the data file
)

// DB is the interface for the database
type DB interface {
	LoadProducts() ([]models.Product, error)
	SaveProducts([]models.Product) error
	GetNextID([]models.Product) int
}

// Client is the implementation of the DB interface
type Client struct{}

// LoadProducts reads products from the data.json file
func (c *Client) LoadProducts() ([]models.Product, error) {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(DataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Product{}, nil // Return empty slice if file doesn't exist
		}
		return nil, err
	}

	var products []models.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// SaveProducts writes products to the data.json file
func (c *Client) SaveProducts(products []models.Product) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(DataFile, data, 0644)
}

// GetNextID calculates the next available ID for a new product
func (c *Client) GetNextID(products []models.Product) int {
	maxID := 0
	for _, p := range products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	return maxID + 1
}