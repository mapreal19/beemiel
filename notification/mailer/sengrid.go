package mailer

import (
	"github.com/astaxie/beego"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func newSengridSender() *emailSender {
	return &emailSender{sengridSender}
}

func sengridSender(email Email) error {
	beego.Info("Sending email through Sendgrid... Recipient: ", email.Tos[0])

	fromMail := mail.NewEmail(email.FromName, email.From)
	body := mail.NewContent("text/html", email.Body)

	m := mail.NewV3Mail()
	m.SetFrom(fromMail)
	m.AddAttachment().Subject = email.Subject
	m.AddContent(body)

	p := mail.NewPersonalization()
	p.AddTos(convertMails(email.Tos)...)
	p.AddCCs(convertMails(email.Ccs)...)
	p.AddBCCs(convertMails(email.Bccs)...)
	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(
		sendgridApiKey,
		"/v3/mail/send",
		"https://api.sendgrid.com",
	)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	_, err := sendgrid.API(request)

	return err
}

func convertMails(addresses []string) []*mail.Email {
	mails := make([]*mail.Email, len(addresses))
	for i, addr := range addresses {
		mails[i] = mail.NewEmail("", addr)
	}
	return mails
}
