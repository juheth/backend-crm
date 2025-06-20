package common

type PaginationDto struct {
	Page  int         `json:"page" validate:"required,number"`
	Limit int         `json:"limit" validate:"required,number"`
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}
