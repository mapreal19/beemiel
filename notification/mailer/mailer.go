package mailer

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/mapreal19/beemiel/envs"
)

type Email struct {
	Subject, Body, From string
	Tos, Ccs, Bccs      []string
	FromName            string //Name of sender ej: M. Bison
}

type emailSender struct {
	send recorderFunc
}

type recorderFunc func(Email) error

var recorder *Email
var sender *emailSender
var mock bool
var sendgridApiKey string

func Init(key string) {

	if envs.IsProduction() {
		sendgridApiKey = key
		sender = newSengridSender()
	} else {
		var send recorderFunc
		send, recorder = mockSend(nil)
		sender = &emailSender{send: send}
	}
}

func Send(email Email) error {
	return sender.Send(email)
}

func GetMock() *Email {
	return recorder
}

func (e *emailSender) Send(email Email) error {

	if !envs.IsProduction() {
		message := fmt.Sprintf(
			"Sending Email from: %v, to: %v, subject: %v",
			email.From,
			email.Tos[0],
			email.Subject,
		)
		beego.Info(message)
	}

	return e.send(email)
}