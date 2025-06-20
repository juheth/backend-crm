package entities

import "time"

type Client struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedBy int       `json:"created_by"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
