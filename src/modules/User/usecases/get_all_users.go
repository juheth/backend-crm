package usecases

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	repository "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
)

type GetAllUsers struct {
	repo repository.UserRepository
}

func NewGetAllUsers(repo repository.UserRepository) *GetAllUsers {
	return &GetAllUsers{
		repo: repo,
	}
}

func (u *GetAllUsers) GetAllUsers() ([]entities.User, error) {
	users, err := u.repo.GetAllUsers(1, 10)
	return users, err
}
