package usecases

import (
	"errors"
	"time"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/utils"
)

type UpdateProduct struct {
	repo repository.ProductRepository
}

func NewUpdateProduct(repo *dao.MySQLProductDao) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}

func (uc *UpdateProduct) Execute(id int, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	if err := utils.ValidateUpdateProduct(req); err != nil {
		return nil, err
	}

	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("producto no encontrado")
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.UpdatedAt = time.Now()

	if err := uc.repo.UpdateProduct(product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		Status:    product.Status,
		CreatedAt: product.CreatedAt,
	}, nil
}
