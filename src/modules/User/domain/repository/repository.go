package repository

import entities "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"

type UserRepository interface {
	GetAllUsers(page, limit int) ([]entities.User, error)
}
