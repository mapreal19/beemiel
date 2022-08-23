package mailer

import (
	"fmt"

	"github.com/astaxie/beego"
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

type sendInterface func(Email)

var recorder *Email
var sender *emailSender
var apiKey string

func Init(key, provider string) {

	if envs.IsProduction() {
		apiKey = key
		switch provider {
		case Provider.MailGun:
			sender = newMailgunSender()
		case Provider.SendGrid:
			sender = newSengridSender()
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
		beego.Info(message)
	}
	e.send(email)
	return nil
}
