package repository

import entities "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"

type UserRepository interface {
	GetAllUsers(page, pageSize int) ([]entities.User, error)
	GetUserById(id int) (entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	FindByEmail(email string) (entities.User, error)
	UpdateUserPasswordByEmail(email string, newPassword string) error
}
