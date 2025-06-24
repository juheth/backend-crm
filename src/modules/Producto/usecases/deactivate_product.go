package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
)

type DeactivateProduct struct {
	repo repository.ProductRepository
}

func NewDeactivateProduct(repo *dao.MySQLProductDao) *DeactivateProduct {
	return &DeactivateProduct{repo: repo}
}

func (uc *DeactivateProduct) Execute(id int) error {
	product, err := uc.repo.GetProductByID(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("producto no encontrado")
	}
	return uc.repo.DeactivateProduct(id)
}
