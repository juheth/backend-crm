package usecases

import (
	"errors"

	auth "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/auth"
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/utils"
)

type LoginUser struct {
	repo repository.UserRepository
}

func NewLoginUser(repo *dao.MySQLUserDao) *LoginUser {
	return &LoginUser{repo: repo}
}

func (l *LoginUser) Execute(payload dto.LoginDTO) (string, *entities.User, error) {
	user, err := l.repo.FindByEmail(payload.Email)
	if err != nil {
		return "", nil, errors.New("usuario no encontrado")
	}

	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		return "", nil, errors.New("contraseña inválida")
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", nil, errors.New("error generando el token")
	}

	user.Password = ""

	return token, &user, nil
}
