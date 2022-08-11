package util

import (
	errorConstants "coda-api/src/constants"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMailBySendGrid SendGridを使用したメール送信
func SendMailBySendGrid(
	emailTo string, subject string, plainTextContent string, htmlContent string,
) error {
	from := mail.NewEmail("Coda", "send_mail@coda.com")
	to := mail.NewEmail(emailTo, emailTo)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return errorConstants.ErrInternalServer
	} else if response.StatusCode == 200 || response.StatusCode == 202 {
		return nil
	} else {
		return errorConstants.ErrInternalServer
	}
}
