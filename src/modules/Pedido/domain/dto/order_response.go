package dto

import "time"

type OrderItemResponse struct {
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unitPrice"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderResponse struct {
	ID        int                 `json:"id"`
	ClientID  int                 `json:"clientId"`
	UserID    int                 `json:"userId"`
	Status    string              `json:"status"`
	Total     float64             `json:"total"`
	OrderDate time.Time           `json:"orderDate"`
	Items     []OrderItemResponse `json:"items"`
}
