package mail

import (
	"fmt"
	"os"
	"github.com/alexamies/cnweb/applog"
	"github.com/alexamies/cnweb/identity"
	"github.com/alexamies/cnweb/webconfig"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendPasswordReset(toUser identity.UserInfo, token string) error {
	fromEmail := webconfig.GetVar("FromEmail")
	from := mail.NewEmail("Do Not Reply", fromEmail)
	subject := "Password Reset"
	to := mail.NewEmail(toUser.FullName, toUser.Email)
	passwordResetURL := webconfig.GetPasswordResetURL()
	plainText := "To reset your password, please go to %s?token=%s . Your username is %s."
	plainTextContent := fmt.Sprintf(plainText, passwordResetURL, token, toUser.UserName)
	htmlText := "<p>To reset your password, please click <a href='%s?token=%s'>here</a>. Your username is %s.</p>"
	htmlContent := fmt.Sprintf(htmlText, passwordResetURL, token, toUser.UserName)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		applog.Info("SendPasswordReset: error, ", err)
		return err
	} else {
		applog.Info("SendPasswordReset: sent email code:, url:",
			response.StatusCode, passwordResetURL)
	}
	return nil
}