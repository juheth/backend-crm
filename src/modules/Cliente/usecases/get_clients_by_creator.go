package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/repository"
)

type GetClientsByCreator struct {
	repo repository.ClientRepository
}

func NewGetClientsByCreator(repo *dao.MySQLClientDao) *GetClientsByCreator {
	return &GetClientsByCreator{
		repo: repo,
	}
}

func (uc *GetClientsByCreator) Execute(creatorID int) ([]entities.Client, error) {
	clients, err := uc.repo.GetClientsByCreator(creatorID)
	if err != nil {
		return nil, err
	}
	if len(clients) == 0 {
		return nil, errors.New("No existen clientes para este usuario")
	}
	return clients, nil
}
