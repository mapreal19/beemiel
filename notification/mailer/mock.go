package mailer

func mockSender() *emailSender {
	var send sendInterface
	send, recorder = mockSend(nil)
	return &emailSender{send}

}

// Inspired by http://tmichel.github.io/2014/10/12/golang-send-test-email/
func mockSend(errToReturn error) (sendInterface, *Email) {
	r := new(Email)

	return func(email Email) error {
		*r = Email{email.Subject, email.Body, email.From, email.Tos, email.Ccs, email.Bccs, email.FromName, email.Attachements}
		return errToReturn
	}, r
}
