package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/repository"
)

type GetClientByID struct {
	repo repository.ClientRepository
}

func NewGetClientById(repo *dao.MySQLClientDao) *GetClientByID {
	return &GetClientByID{repo: repo}
}

func (uc *GetClientByID) Execute(id int) (*entities.Client, error) {
	client, err := uc.repo.GetClientByID(id)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, errors.New("El cliente no existe")
	}
	return client, nil
}
