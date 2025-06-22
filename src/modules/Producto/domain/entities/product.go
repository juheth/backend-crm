package entities

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}
