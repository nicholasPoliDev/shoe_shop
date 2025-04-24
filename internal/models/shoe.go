package models

type Shoe struct {
	ID          int     `json:"id"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
