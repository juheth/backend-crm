package User

import (
	dto "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
)

func WelcomeEmail(password, email string) error {

	send := dto.SendEmail{
		Subject:   "Bienvenido a Nexcol",
		Recipient: email,
		PlainText: "Bienvenido!",
		Template:  "src/modules/user/utils/mailgun/templates/welcome_email/welcome_email.html",
	}

	data := map[string]interface{}{
		"NewPassword": password,
	}

	return SendEmail(send, data)
}
