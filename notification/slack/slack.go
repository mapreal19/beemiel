package slack

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/mapreal19/beemiel/v2/envs"
	"github.com/slack-go/slack"
)

type Message struct {
	Channel, Body, Header string
	IconURL, Username     string
	Attachments           []Attachment
}

type Attachment struct {
	Data       []byte
	Name, Type string
}

type slackSender struct {
	send sendInterface
}
type sendInterface func(Message) error

var slackApiKey string
var recorder *Message
var sender *slackSender

func Init(key string) {

	if envs.IsProduction() {
		slackApiKey = key
		sender = newSlackSender()
	} else {
		sender = mockSender()
	}
}

func GetAPI() *slack.Client {
	return slack.New(slackApiKey)
}

func Send(message Message) error {
	return sender.Send(message)
}

func GetMock() *Message {
	return recorder
}

func (e *slackSender) Send(message Message) error {

	if !envs.IsProduction() {
		text := fmt.Sprintf(
			"Sending Slack from: %v, to: %v, subject: %v",
			message.Username,
			message.Channel,
			message.Header,
		)
		logs.Info(text)
	}

	return e.send(message)
}
