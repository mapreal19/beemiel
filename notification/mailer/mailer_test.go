package mailer_test

import (
	"fmt"
	"os"

	"github.com/mapreal19/beemiel/envs"
	"github.com/mapreal19/beemiel/notification/mailer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mailer", func() {
	Describe("Mock email", func() {
		It("get same as sended", func() {
			envs.Init("")
			_ = os.Setenv("BEEGO_ENV", "test")

			mailer.Init("dontNeeded")
			var email mailer.Email
			email.From = "m.bison@bad.bad"
			email.FromName = "M. Bison"
			email.Subject = "Surrender!"
			email.Body = "I go to win!"
			email.Tos = []string{"ryu@good.good"}
			email.Ccs = []string{"blanka@good.good"}
			email.Bccs = []string{"chun.li@good.good"}

			err := mailer.Send(email)
			Expect(err).To(BeNil())

			mockEmail := mailer.GetMock()
			Expect(mockEmail.From).To(Equal("m.bison@bad.bad"))
			Expect(mockEmail.FromName).To(Equal("M. Bison"))
			Expect(mockEmail.Subject).To(Equal("Surrender!"))
			Expect(mockEmail.Body).To(Equal("I go to win!"))
			Expect(len(mockEmail.Tos)).To(Equal(1))
			Expect(len(mockEmail.Ccs)).To(Equal(1))
			Expect(len(mockEmail.Bccs)).To(Equal(1))
		})
	})

	Describe("Mock email file", func() {
		It("get attachment", func() {
			envs.Init("")
			_ = os.Setenv("BEEGO_ENV", "test")
			mailer.Init("dontNeeded")

			var attachement mailer.Attachment
			attachement.Name = "hello.txt"
			attachement.Type = "text/plain"
			attachement.Data = []byte("Hello world")

			var email mailer.Email
			email.From = "m.bison@bad.bad"
			email.Tos = []string{"ryu@good.good"}
			email.Attachments = []mailer.Attachment{attachement}

			err := mailer.Send(email)
			Expect(err).To(BeNil())

			mockEmail := mailer.GetMock()
			Expect(len(mockEmail.Attachments)).To(Equal(1))
			Expect(mockEmail.Attachments[0].Name).To(Equal("hello.txt"))
			Expect(mockEmail.Attachments[0].Type).To(Equal("text/plain"))
			Expect(mockEmail.Attachments[0].Data).To(Equal([]byte("Hello world")))
		})
	})
	Describe("1 To at least", func() {
		It("get same file parameters", func() {
			envs.Init("")
			_ = os.Setenv("BEEGO_ENV", "test")
			mailer.Init("dontNeeded")

			var email mailer.Email
			email.From = "m.bison@bad.bad"

			err := mailer.Send(email)
			Expect(err).To(Equal(fmt.Errorf("There must be at least 1 'To' recipient")))

			mockEmail := mailer.GetMock()
			Expect(mockEmail.From).To(Equal(""))
		})
	})
})
