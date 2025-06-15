package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
)

type UpdateUser struct {
	repo repository.UserRepository
}

func NewUpdateUser(repo *dao.MySQLUserDao) *UpdateUser {
	return &UpdateUser{
		repo: repo,
	}
}

func (u *UpdateUser) Execute(user *entities.User) error {
	return u.repo.UpdateUser(user)
}
