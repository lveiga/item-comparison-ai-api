package repositories

import (
	"encoding/json"
	"item-comparison-ai-api/config"
	"item-comparison-ai-api/internal/database"
	"os"
	"sync"
)

type BaseRepositoryInterface interface {
	Load() ([]byte, error)
	Save([]interface{}) error
}

// Client is the implementation of the DB interface
type Client struct {
	FileStore database.FileStore
	config    *config.AppConfig
}

var (
	mu sync.Mutex // Mutex to protect access to the data file
)

var dataFilePath = os.Getenv("DATA_FILE_PATH")

// Load reads  from the data.json file
func (c *Client) Load() ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()

	data, err := c.FileStore.Read(c.config.DatabasePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // Return empty slice if file doesn't exist
		}
		return nil, err
	}

	return data, nil
}

// Save writes  to the data.json file
func (c *Client) Save(model []interface{}) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return err
	}

	return c.FileStore.Write(c.config.DatabasePath, data, 0644)
}

func NewBaseRepository(fileStore database.FileStore, conf *config.AppConfig) *Client {
	return &Client{
		FileStore: fileStore,
		config:    conf,
	}
}
