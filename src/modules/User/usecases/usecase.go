package usecases

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
)

type FindAllUser struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *FindAllUser {
	return &FindAllUser{
		repo: repo,
	}
}

func (u *FindAllUser) FindAllUser() {
}
