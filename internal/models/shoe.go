package models

import "fmt"

// Shoe represents a product in the inventory system.
type Shoe struct {
	ID          int    `json:"id"`
	Brand       string `json:"brand"` // e.g., Nike, Puma, Asics
	Description string `json:"description"`
	Category    string `json:"category"` // e.g., Running, Casual, Basketball
	Price       int    `json:"price"`    // Price in cents to avoid float precision
	Stock       int    `json:"stock"`    // Quantity available in stock
}

// NewShoe is a constructor function that creates and return a new Shoe instance.
// It ensures that a valid Shoe struct is always returned with no nil fields.
// It also validates that the price and stock are non-negative and that the brand and
// category are not empty.
func NewShoe(id int, brand, description, category string, price, stock int) (*Shoe, error) {
	// price and stock validation
	if price < 0 {
		return nil, fmt.Errorf("price cannot be negative")
	}
	if stock < 0 {
		return nil, fmt.Errorf("stock cannot be negative")
	}
	//brand and category validation
	if brand == "" {
		return nil, fmt.Errorf("brand cannot be empty")
	}
	if category == "" {
		return nil, fmt.Errorf("category cannot be empty")
	}

	// Return the new Shoe instance if all validations pass
	return &Shoe{
		ID:          id,
		Brand:       brand,
		Description: description,
		Category:    category,
		Price:       price,
		Stock:       stock,
	}, nil
}
