package mailer

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/mapreal19/beemiel/envs"
)

var Provider = struct {
	MailGun, SendGrid string
}{
	MailGun:  "MAILGUN",
	SendGrid: "SENDGRID",
}

type Email struct {
	Subject, Body, PlainBody, From string
	Tos, Ccs, Bccs                 []string
	FromName                       string //Name of sender ej: M. Bison
	Attachments                    []Attachment
}

type Attachment struct {
	Data       []byte
	Name, Type string
}

type emailSender struct {
	send sendInterface
}

type sendInterface func(Email) error

var recorder *Email
var sender *emailSender

type globalConf struct {
	ApiKey string
}

func Init(key, provider string) {
	if envs.IsProduction() {

		g := globalConf{
			ApiKey: key,
		}
		switch provider {
		case Provider.MailGun:
			sender = newMailgunSender(g)
		case Provider.SendGrid:
			sender = newSengridSender(g)
		}

	} else {
		sender = mockSender()
	}
}

func Send(email Email) error {
	return sender.Send(email)
}

func GetMock() *Email {
	return recorder
}

func (e *emailSender) Send(email Email) error {
	if len(email.Tos) < 1 {
		return fmt.Errorf("There must be at least 1 'To' recipient")
	}

	if !envs.IsProduction() {
		message := fmt.Sprintf(
			"Sending Email from: %v, to: %v, subject: %v",
			email.From,
			email.Tos[0],
			email.Subject,
		)
		logs.Info(message)
	}
	return e.send(email)
}
