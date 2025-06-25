package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
)

type GetProductByID struct {
	repo repository.ProductRepository
}

func NewGetProductByID(repo *dao.MySQLProductDao) *GetProductByID {
	return &GetProductByID{repo: repo}
}

func (uc *GetProductByID) Execute(id int) (*dto.ProductResponse, error) {
	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("producto no encontrado")
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
