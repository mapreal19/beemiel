package slack

import "github.com/slack-go/slack"

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

	return err
}
