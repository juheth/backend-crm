package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
)

type ActivateProduct struct {
	repo repository.ProductRepository
}

func NewActivateProduct(repo *dao.MySQLProductDao) *ActivateProduct {
	return &ActivateProduct{repo: repo}
}

func (uc *ActivateProduct) Execute(id int) (*dto.ProductResponse, error) {
	product, err := uc.repo.GetProductByIDAnyStatus(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("producto no encontrado")
	}
	if product.Status {
		return nil, errors.New("el producto ya está activo")
	}
	if err := uc.repo.ActivateProduct(id); err != nil {
		return nil, err
	}
	product.Status = true
	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}
