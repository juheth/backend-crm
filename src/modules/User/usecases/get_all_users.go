package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
)

type GetAllUsers struct {
	repo repository.UserRepository
}

func NewGetAllUsers(repo *dao.MySQLUserDao) *GetAllUsers {
	return &GetAllUsers{
		repo: repo,
	}
}

func (u *GetAllUsers) Execute() ([]entities.User, error) {
	users, err := u.repo.GetAllUsers(1, 10)
	return users, err
}
