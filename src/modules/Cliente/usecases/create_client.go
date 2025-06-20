package usecases

import (
	"errors"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/repository"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
)

type CreateClient struct {
	repo repository.ClientRepository
}

func NewCreateClient(repo *dao.MySQLClientDao) *CreateClient {
	return &CreateClient{
		repo: repo,
	}
}

func (c *CreateClient) Execute(client *entities.Client) error {
	exists, err := c.repo.ExistsByEmailOrPhone(client.Email, client.Phone)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Ya existe un cliente con ese email o teléfono")
	}
	return c.repo.CreateClient(client)
}
