package mailer

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/mapreal19/beemiel/envs"
)

type Email struct {
	Subject, Body, PlainBody, From string
	Tos, Ccs, Bccs                 []string
	FromName                       string //Name of sender ej: M. Bison
	Attachements                   []Attachement
}

type Attachement struct {
	Data       []byte
	Name, Type string
}

type emailSender struct {
	send sendInterface
}

type sendInterface func(Email)

var recorder *Email
var sender *emailSender
var sendgridApiKey string

func Init(key string) {

	if envs.IsProduction() {
		sendgridApiKey = key
		sender = newSengridSender()
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
