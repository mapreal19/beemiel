package mailer

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/mailgun/mailgun-go"
)

func newMailgunSender() *emailSender {
	return &emailSender{mailgunSender}
}

func mailgunSender(email Email) {
	beego.Info("Sending email through MailGun... Recipient: ", email.Tos[0])
	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), apiKey)

	message := mg.NewMessage(
		fmt.Sprint("%s<%s>", email.FromName, email.From),
		email.Subject,
		email.PlainBody,
		email.Tos...)

	mailGunAddCCs(message, email.Ccs)
	mailGunAddBCCs(message, email.Bccs)
	mailGunAddAttachments(message, email.Attachments)
	message.SetHtml(email.Body)

	mg.Send(message)

}
func mailGunAddAttachments(message *mailgun.Message, attachment []Attachment) {
	for _, att := range attachment {
		message.AddBufferAttachment(att.Name, att.Data)
	}
}

func mailGunAddCCs(message *mailgun.Message, addresses []string) {
	for _, addr := range addresses {
		message.AddCC(addr)
	}
}

func mailGunAddBCCs(message *mailgun.Message, addresses []string) {
	for _, addr := range addresses {
		message.AddBCC(addr)
	}
}
