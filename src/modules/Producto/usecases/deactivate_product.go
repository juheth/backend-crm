package usecases

import (
    "errors"

    dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/repository"
)

type DeactivateProduct struct {
    repo repository.ProductRepository
}

func NewDeactivateProduct(repo *dao.MySQLProductDao) *DeactivateProduct {
    return &DeactivateProduct{repo: repo}
}

func (uc *DeactivateProduct) Execute(id int) (*dto.ProductResponse, error) {
    product, err := uc.repo.GetProductByID(id)
    if err != nil {
        return nil, err
    }
    if product == nil {
        return nil, errors.New("producto no encontrado")
    }

    if err := uc.repo.DeactivateProduct(id); err != nil {
        return nil, err
    }

    product.Status = false

    return &dto.ProductResponse{
        ID:        product.ID,
        Name:      product.Name,
        Price:     product.Price,
        Stock:     product.Stock,
        Status:    product.Status,
        CreatedAt: product.CreatedAt,
    }, nil
}