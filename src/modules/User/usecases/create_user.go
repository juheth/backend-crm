package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
	"golang.org/x/crypto/bcrypt"
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
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return u.repo.CreateUser(user)
}
