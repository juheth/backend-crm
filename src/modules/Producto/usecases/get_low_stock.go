package usecases

import (
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
    dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
)

type GetLowStockProducts struct {
    repo *dao.MySQLProductDao
}

func NewGetLowStockProducts(repo *dao.MySQLProductDao) *GetLowStockProducts {
    return &GetLowStockProducts{repo: repo}
}

func (uc *GetLowStockProducts) Execute(threshold int) ([]dto.ProductResponse, error) {
    products, err := uc.repo.GetLowStock(threshold)
    if err != nil {
        return nil, err
    }
    var result []dto.ProductResponse
    for _, p := range products {
        result = append(result, dto.ProductResponse{
            ID:          p.ID,
            Name:        p.Name,
            Description: p.Description,
            Price:       p.Price,
            Stock:       p.Stock,
            Status:      p.Status,
            CreatedAt:   p.CreatedAt,
            UpdatedAt:   p.UpdatedAt,
        })
    }
    return result, nil
}