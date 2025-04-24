package models

import "testing"

func TestNewShoe(t *testing.T) {
	tests := []struct {
		name        string
		id          int
		brand       string
		description string
		category    string
		price       int
		stock       int
		expectedErr error
	}{
		// Test for a invalid ID.
		// Currently, the ID is not provided
		// because in future database integration, the ID will be auto-incremented.
		// We don't need to validate it manually in this test case.

		// Test for a valid case.
		// We expect no error because this case is valid
		{
			name:        "Valid Shoe",
			id:          1,
			brand:       "Asics",
			description: "A shoe perfect to run everywhere",
			category:    "running",
			price:       5000, //price in cents
			stock:       10,
			expectedErr: nil,
		},
		// Test for negative price case.
		// We expect an error because a positive price is required
		{
			name:        "Invalid Price",
			id:          2,
			brand:       "Mitzuno",
			description: "A shoe adapted for intensive sports",
			category:    "indoor sports",
			price:       -100,
			stock:       7,
			expectedErr: ErrInvalidPrice,
		},
		// Test for negative stock case.
		// We expect an error because a positive quantity is required
		{
			name:        "Invalid Stock",
			id:          3,
			brand:       "Asics",
			description: "Basketball shoe",
			category:    "Basketball",
			price:       6000,
			stock:       -1,
			expectedErr: ErrInvalidStock,
		},
		// Test for empty string brand case.
		// We expect an error because a valid brand is required
		{
			name:        "Empty Brand",
			id:          4,
			brand:       "",
			description: "Hiking shoe",
			category:    "Hiking",
			price:       7000,
			stock:       20,
			expectedErr: ErrEmptyBrand,
		},
		// Test for empty string category
		{
			name:        "Empty Category",
			id:          5,
			brand:       "Mizuno",
			description: "Training shoe",
			category:    "",
			price:       8000,
			stock:       30,
			expectedErr: ErrEmptyCategory,
		},
		// Test for price equal 0 case.
		// We expect no error because a shoe could be in promo or given free
		{
			name:        "Zero Price",
			id:          6,
			brand:       "Adidas",
			description: "Running shoe",
			category:    "Running",
			price:       0,
			stock:       10,
			expectedErr: nil,
		},
		// Test for stock equal 0 case.
		// We expect no error because a shoe could be sold out
		{
			name:        "Zero Stock",
			id:          7,
			brand:       "Reebok",
			description: "Casual shoe",
			category:    "Casual",
			price:       4000,
			stock:       0,
			expectedErr: nil,
		},
		// test for price and stock limit values
		// Testing edge cases with maimum integer value to ensure that
		// the system handles boundary cases correctly and that it does not
		// overflow or cause unexpected behavior.
		{
			name:        "Max Values",
			id:          8,
			brand:       "Nike",
			description: "Sport shoe",
			category:    "Sport",
			price:       2147483647, // max limit for integer
			stock:       2147483647, // max limit for integer
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// A new instance of Shoe
			shoe, err := NewShoe(tt.id, tt.brand, tt.description, tt.category, tt.price, tt.stock)

			// check if the returned error is what attended
			if err != nil && err != tt.expectedErr {
				t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
			}

			// if there is no error, proceed with shoe's value verification.
			if err == nil && shoe == nil {
				t.Fatalf("expected shoe to be created, got nil")
			}

			// if the shoe has been created, verify data consistency
			if shoe != nil {
				if shoe.ID != tt.id {
					t.Errorf("expected ID %v, got %v", tt.id, shoe.ID)
				}
				if shoe.Brand != tt.brand {
					t.Errorf("expected brand %v, got %v", tt.brand, shoe.Brand)
				}
				if shoe.Description != tt.description {
					t.Errorf("expected description %v, got %v", tt.description, shoe.Description)
				}
				if shoe.Category != tt.category {
					t.Errorf("expected category %v, got %v", tt.category, shoe.Category)
				}
				if shoe.Price != tt.price {
					t.Errorf("expected price %v, got %v", tt.price, shoe.Price)
				}
				if shoe.Stock != tt.stock {
					t.Errorf("expected stock %v, got %v", tt.stock, shoe.Stock)
				}
			}
		})
	}
}
