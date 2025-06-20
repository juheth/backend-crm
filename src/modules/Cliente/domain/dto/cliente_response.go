package dto

import "time"

type ClientResponseDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"createdAt"`
}
