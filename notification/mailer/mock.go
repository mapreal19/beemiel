package mailer

// Inspired by http://tmichel.github.io/2014/10/12/golang-send-test-email/
func mockSend(errToReturn error) (recorderFunc, *Email) {
	r := new(Email)

	return func(email Email) error {
		*r = Email{email.Subject, email.Body, email.From, email.Tos, email.Ccs, email.Bccs, email.FromName}
		return errToReturn
	}, r
}
