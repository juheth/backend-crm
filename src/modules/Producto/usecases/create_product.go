package usecases

import (
	"errors"
	"time"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/utils"
)

type CreateProduct struct {
	repo repository.ProductRepository
}

func NewCreateProduct(repo *dao.MySQLProductDao) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (uc *CreateProduct) Execute(request dto.CreateProductRequest) (*dto.ProductResponse, error) {
	if err := utils.ValidateCreateProduct(request); err != nil {
		return nil, err
	}

	exists, err := uc.repo.ExistsByName(request.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("ya existe un producto con ese nombre")
	}

	statusBool, err := utils.ParseStatus(request.Status)
	if err != nil {
		return nil, err
	}

	product := &entities.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Status:      statusBool,
		CreatedAt:   time.Now(),
	}

	if err := uc.repo.Create(product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
	}, nil
}
