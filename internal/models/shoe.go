package models

import "fmt"

// Predefined validation errors for Shoe fields.
var (
	ErrInvalidPrice  = fmt.Errorf("price cannot be negative")
	ErrInvalidStock  = fmt.Errorf("stock cannot be negative")
	ErrEmptyBrand    = fmt.Errorf("brand cannot be empty")
	ErrEmptyCategory = fmt.Errorf("category cannot be empty")
)

// Shoe represents a product in the inventory system.
//
// It includes details such as brand, description, category, price, and stock.
// The Price is stored in cents to avoid floating point precision issues.
// This struct is used in CRUD operations and can be serialized to/from JSON.
type Shoe struct {
	ID          int    `json:"id"`          // Unique identifier for the shoe.
	Brand       string `json:"brand"`       // Brand name (e.g., Nike, Puma, Asics).
	Description string `json:"description"` // Description (e.g., color, material).
	Category    string `json:"category"`    // Shoe category (e.g., Running, Casual).
	Price       int    `json:"price"`       // Price in cents.
	Stock       int    `json:"stock"`       // Available quantity in stock.
}

// NewShoe creates a new Shoe instance after validating the input data.
//
// It returns an error if any of the fields are invalid (e.g., negative price,
// empty brand or category). On success, it returns a pointer to a valid Shoe.
func NewShoe(id int, brand, description, category string, price, stock int) (*Shoe, error) {
	if err := validateShoeData(brand, category, price, stock); err != nil {
		return nil, err
	}

	return &Shoe{
		ID:          id,
		Brand:       brand,
		Description: description,
		Category:    category,
		Price:       price,
		Stock:       stock,
	}, nil
}

// validateShoeData checks that the brand and category are not empty,
// and that price and stock are non-negative.
//
// It returns an error if any validation rule fails.
func validateShoeData(brand, category string, price, stock int) error {
	if price < 0 {
		return ErrInvalidPrice
	}
	if stock < 0 {
		return ErrInvalidStock
	}
	if brand == "" {
		return ErrEmptyBrand
	}
	if category == "" {
		return ErrEmptyCategory
	}
	return nil
}
