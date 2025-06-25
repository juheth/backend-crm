package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
)

type GetAllProducts struct {
	repo *dao.MySQLProductDao
}

func NewGetAllProducts(repo *dao.MySQLProductDao) *GetAllProducts {
	return &GetAllProducts{repo: repo}
}

func (uc *GetAllProducts) Execute() ([]dto.ProductResponse, error) {
	products, err := uc.repo.GetAll()
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
		})
	}
	return result, nil
}
