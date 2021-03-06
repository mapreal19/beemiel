package mailer

import (
	"encoding/base64"

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

	m := mail.NewV3Mail()
	m.SetFrom(fromMail)
	contents := buildBodies(email)
	m.AddContent(contents...)

	for _, atachement := range email.Attachements {
		encoded := base64.StdEncoding.EncodeToString(atachement.Data)
		a := mail.NewAttachment()
		a.SetContent(encoded)
		a.SetType(atachement.Type)
		a.SetFilename(atachement.Name)
		a.SetDisposition("attachment")
		m.AddAttachment(a)
	}

	p := mail.NewPersonalization()
	p.AddTos(convertMails(email.Tos)...)
	p.AddCCs(convertMails(email.Ccs)...)
	p.AddBCCs(convertMails(email.Bccs)...)
	p.Subject = email.Subject

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

func buildBodies(email Email) (contents []*mail.Content) {
	if email.PlainBody != "" {
		contents = append(contents, mail.NewContent("text/plain", email.PlainBody))
	}
	if email.Body != "" {
		contents = append(contents, mail.NewContent("text/html", email.Body))
	}
	return
}
