package mailer

func mockSender() *emailSender {
	var send sendInterface
	send, recorder = mockSend()
	return &emailSender{send}

}

// Inspired by http://tmichel.github.io/2014/10/12/golang-send-test-email/
func mockSend() (sendInterface, *Email) {
	r := new(Email)

	return func(email Email) {
		*r = Email{email.Subject, email.Body, email.PlainBody, email.From, email.Tos, email.Ccs, email.Bccs, email.FromName, email.Attachments}
	}, r
}
