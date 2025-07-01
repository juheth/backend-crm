package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginResponse struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}
