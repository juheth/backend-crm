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

func (uc *CreateProduct) Execute(Request dto.CreateProductRequest) (*dto.ProductResponse, error) {
	if err := utils.ValidateProductInput(Request.Name, Request.Description, Request.Price, Request.Stock); err != nil {
		return nil, err
	}

	exists, err := uc.repo.ExistsByName(Request.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("ya existe un producto con ese nombre")
	}

	product := &entities.Product{
		Name:        Request.Name,
		Description: Request.Description,
		Price:       Request.Price,
		Stock:       Request.Stock,
		Status:      true,
		CreatedAt:   time.Now(),
	}

	err = uc.repo.Create(product)
	if err != nil {
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
