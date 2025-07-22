package User

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"os"
	"time"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
)

func SendEmail(send dto.SendEmail, data interface{}) error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	yourDomain := os.Getenv("MAILGUN_DOMAIN")
	privateAPIKey := os.Getenv("MAILGUN_SECRET_KEY")
	if yourDomain == "" || privateAPIKey == "" {
		return errors.New("MAILGUN_DOMAIN or MAILGUN_SECRET_KEY not set in .env file")
	}

	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	sender := "Nexcol <no-reply@" + yourDomain + ">"
	subject := send.Subject
	recipient := send.Recipient
	plainText := send.PlainText

	// Parsear plantilla HTML
	tmpl, err := template.ParseFiles("src/modules/user/utils/mailgun/templates/welcome_email/welcome_email.html")
	if err != nil {
		return err
	}

	var htmlBody bytes.Buffer
	if err := tmpl.Execute(&htmlBody, data); err != nil {
		return err
	}

	message := mailgun.NewMessage(sender, subject, plainText, recipient)
	message.SetHtml(htmlBody.String())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Enviar el correo
	_, _, err = mg.Send(ctx, message)
	return err
}
