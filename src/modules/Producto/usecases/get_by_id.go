package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
)

type GetProductByID struct {
	repo repository.ProductRepository
}

func NewGetProductByID(repo *dao.MySQLProductDao) *GetProductByID {
	return &GetProductByID{repo: repo}
}

func (uc *GetProductByID) Execute(id int) (*entities.Product, error) {
	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("producto no encontrado")
	}
	return product, nil

}
