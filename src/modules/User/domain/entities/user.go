package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Email     string         `json:"email" gorm:"uniqueIndex"`
	Status    string         `json:"status" gorm:"default:'active'"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
