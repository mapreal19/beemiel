package mailer

import (
	"encoding/base64"

	"github.com/beego/beego/v2/core/logs"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGrid struct {
	globalConf
	Region, Domain string
}

var sendGridConf sendGrid

func newSengridSender(g globalConf) *emailSender {
	sendGridConf.globalConf = g
	return &emailSender{sendGridConf.sengridSender}
}

func (s *sendGrid) sengridSender(email Email) error {
	logs.Info("Sending email through Sendgrid... Recipient: ", email.Tos[0])

	fromMail := mail.NewEmail(email.FromName, email.From)

	m := mail.NewV3Mail()
	m.SetFrom(fromMail)
	contents := s.buildBodies(email)
	m.AddContent(contents...)

	for _, attachment := range email.Attachments {
		encoded := base64.StdEncoding.EncodeToString(attachment.Data)
		a := mail.NewAttachment()
		a.SetContent(encoded)
		a.SetType(attachment.Type)
		a.SetFilename(attachment.Name)
		a.SetDisposition("attachment")
		m.AddAttachment(a)
	}

	p := mail.NewPersonalization()
	p.AddTos(s.convertMails(email.Tos)...)
	p.AddCCs(s.convertMails(email.Ccs)...)
	p.AddBCCs(s.convertMails(email.Bccs)...)
	p.Subject = email.Subject

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(
		s.ApiKey,
		"/v3/mail/send",
		"https://api.sendgrid.com",
	)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	sendgrid.MakeRequestAsync(request)
	return nil
}

func (s *sendGrid) convertMails(addresses []string) []*mail.Email {
	mails := make([]*mail.Email, len(addresses))
	for i, addr := range addresses {
		mails[i] = mail.NewEmail("", addr)
	}
	return mails
}

func (s *sendGrid) buildBodies(email Email) (contents []*mail.Content) {
	if email.PlainBody != "" {
		contents = append(contents, mail.NewContent("text/plain", email.PlainBody))
	}
	if email.Body != "" {
		contents = append(contents, mail.NewContent("text/html", email.Body))
	}
	return
}
