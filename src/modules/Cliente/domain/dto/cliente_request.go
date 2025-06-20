package dto

type CreateClientRequestDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedBy int    `json:"created_by"`
}
