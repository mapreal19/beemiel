package slack

import (
	"github.com/slack-go/slack"
)

func newSlackSender() *slackSender {
	return &slackSender{send}
}

func send(message Message) error {

	api := slack.New(slackApiKey)

	_, _, err := api.PostMessage(
		message.Channel,
		slack.MsgOptionText(message.Body, false),
		slack.MsgOptionAsUser(false),
		slack.MsgOptionUsername(message.Username),
		slack.MsgOptionIconURL(message.IconURL),
	)

	attachFiles(api, message)
	return err
}

func attachFiles(api *slack.Client, message Message) {
	for _, attachment := range message.Attachments {
		params := slack.FileUploadParameters{
			Title:    attachment.Name,
			Filetype: attachment.Type,
			Filename: attachment.Name,
			Content:  string(attachment.Data),
			Channels: []string{message.Channel},
		}
		api.UploadFile(params)
	}
}
