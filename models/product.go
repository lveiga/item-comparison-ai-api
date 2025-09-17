package models

// Product represents the model for a product
type Product struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	ImageURL    string            `json:"image_url"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Rating      float64           `json:"rating"`
	Specifications map[string]string `json:"specifications"`
}
