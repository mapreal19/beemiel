package mailer_test

import (
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

			mailer.Send(email)

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
})
