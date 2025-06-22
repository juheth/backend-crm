package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gte=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}
