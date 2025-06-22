package dto

import "time"

type ProductResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
