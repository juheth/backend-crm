package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
)

type CreateUsers struct {
	repo repository.UserRepository
}

func NewCreateUsers(repo *dao.MySQLUserDao) *CreateUsers {
	return &CreateUsers{
		repo: repo,
	}
}

func (u *CreateUsers) Execute(user *entities.User) error {
	return u.repo.CreateUser(user)

}
