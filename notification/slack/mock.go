package slack

func mockSender() *slackSender {
	var send sendInterface
	send, recorder = mockSend(nil)
	return &slackSender{send}

}

// Inspired by http://tmichel.github.io/2014/10/12/golang-send-test-email/
func mockSend(errToReturn error) (sendInterface, *Message) {
	r := new(Message)

	return func(message Message) error {
		*r = Message{message.Channel, message.Body, message.Header, message.IconURL, message.Username}
		return errToReturn
	}, r
}
