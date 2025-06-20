package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/repository"
)

type GetAllClients struct {
	repo repository.ClientRepository
}

func NewGetAllClients(repo *dao.MySQLClientDao) *GetAllClients {
	return &GetAllClients{repo: repo}
}

func (uc *GetAllClients) Execute() ([]*entities.Client, error) {
	clients, err := uc.repo.GetAllClients()
	if err != nil {
		return nil, err
	}
	if len(clients) == 0 {
		return nil, errors.New("No existen clientes registrados")
	}
	return clients, nil
}
