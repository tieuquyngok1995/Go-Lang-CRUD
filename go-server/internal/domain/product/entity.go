package product

import "time"

type Product struct {
	ID          int64     `db:"id"          json:"id"`
	Name        string    `db:"name"        json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price"       json:"price"`
	Stock       int       `db:"stock"       json:"stock"`
	CreatedAt   time.Time `db:"created_at"  json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"  json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"        binding:"required,min=1"`
	Description string  `json:"description"`
	Price       float64 `json:"price"       binding:"required,min=0"`
	Stock       int     `json:"stock"       binding:"min=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"        binding:"omitempty,min=1"`
	Description string  `json:"description"`
	Price       float64 `json:"price"       binding:"omitempty,min=0"`
	Stock       int     `json:"stock"       binding:"omitempty,min=0"`
}
