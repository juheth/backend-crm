package dto

import "time"

type UserDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
