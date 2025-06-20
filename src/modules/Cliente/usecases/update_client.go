package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/repository"
)

type UpdateClient struct {
	repo repository.ClientRepository
}

func NewUpdateClient(repo *dao.MySQLClientDao) *UpdateClient {
	return &UpdateClient{repo: repo}
}

func (uc *UpdateClient) Execute(client *entities.Client) error {

	existing, err := uc.repo.GetClientByID(client.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("El cliente no existe")
	}
	return uc.repo.UpdateClient(client)
}
