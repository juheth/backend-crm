package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
)

type GetUserById struct {
	repo repository.UserRepository
}

func NewGetUserById(repo *dao.MySQLUserDao) *GetUserById {
	return &GetUserById{
		repo: repo,
	}
}

func (u *GetUserById) GetUserById(id int) (entities.User, error) {
	user, err := u.repo.GetUserById(id)
	return user, err
}
