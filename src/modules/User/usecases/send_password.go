package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/repository"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/utils"
	mailgun "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/utils/mailgun"
)

type SendPassword struct {
	repo repository.UserRepository
}

func NewSendPassword(repo *dao.MySQLUserDao) *SendPassword {
	return &SendPassword{repo: repo}
}

func (uc *SendPassword) Execute(input *dto.SendPasswordRequest) error {
	user, err := uc.repo.FindByEmail(input.Email)
	if err != nil {
		return err
	}

	newPassword := utils.GenerateSecurePassword(10)
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := uc.repo.UpdateUserPasswordByEmail(input.Email, hashedPassword); err != nil {
		return err
	}

	return mailgun.WelcomeEmail(newPassword, user.Email)
}
