package entities

import "time"

type Order struct {
	ID              int         `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientID        int         `json:"clientId"`
	UserID          int         `json:"userId"`
	PaymentMethodID int         `json:"paymentMethodId"`
	Status          string      `json:"status"`
	Total           float64     `json:"total"`
	OrderDate       time.Time   `json:"orderDate" gorm:"autoCreateTime"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        int     `json:"-" gorm:"primaryKey;autoIncrement"`
	OrderID   int     `json:"-"`
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unitPrice"`
	Subtotal  float64 `json:"subtotal" gorm:"->"`
}
