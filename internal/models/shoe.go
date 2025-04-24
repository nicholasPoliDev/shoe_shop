package models

// Shoe represents a product in the inventory system.
type Shoe struct {
	ID          int    `json:"id"`
	Brand       string `json:"brand"` // e.g., Nike, Puma, Asics
	Description string `json:"description"`
	Category    string `json:"category"` // e.g., Running, Casual, Basketball
	Price       int    `json:"price"`    // Price in cents to avoid float precision
	Stock       int    `json:"stock"`    // Quantity available in stock
}
