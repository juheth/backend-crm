package repository

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
)

type ClientRepository interface {
	CreateClient(client *entities.Client) error
	ExistsByEmailOrPhone(email, phone string) (bool, error)
	GetAllClients() ([]*entities.Client, error)
	GetClientByID(id int) (*entities.Client, error)
}
