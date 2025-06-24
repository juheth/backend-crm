package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
)

type DeleteProduct struct {
    repo repository.ProductRepository
}

func NewDeleteProduct(repo *dao.MySQLProductDao) *DeleteProduct {
    return &DeleteProduct{repo: repo}
}

func (uc *DeleteProduct) Execute(id int) error {
    product, err := uc.repo.GetProductByID(id)
    if err != nil {
        return err
    }
    if product == nil {
        return errors.New("producto no encontrado")
    }
    return uc.repo.DeleteProduct(id)
}