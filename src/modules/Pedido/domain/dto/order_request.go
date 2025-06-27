package dto

type OrderItemRequest struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type CreateOrderRequest struct {
	ClientID        int                `json:"clientId"`
	PaymentMethodID int                `json:"paymentMethodId"`
	Items           []OrderItemRequest `json:"items"`
}
