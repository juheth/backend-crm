package entities

import "time"

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
