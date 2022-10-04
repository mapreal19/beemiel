package slack_test

import (
	"os"

	"github.com/mapreal19/beemiel/v2/envs"
	"github.com/mapreal19/beemiel/v2/notification/slack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mailer", func() {
	Describe("Mock email", func() {
		It("get same as sended", func() {
			envs.Init("")
			_ = os.Setenv("BEEGO_ENV", "test")

			slack.Init("dontNeeded")
			var msg slack.Message
			msg.Channel = "abcdef"
			msg.Body = "I go to win!"
			msg.Header = "Surrender!"
			msg.IconURL = "bison.png"
			msg.Username = "M. Bison"

			slack.Send(msg)

			mockMsg := slack.GetMock()
			Expect(mockMsg.IconURL).To(Equal("bison.png"))
			Expect(mockMsg.Username).To(Equal("M. Bison"))
			Expect(mockMsg.Header).To(Equal("Surrender!"))
			Expect(mockMsg.Body).To(Equal("I go to win!"))
			Expect(mockMsg.Channel).To(Equal("abcdef"))
		})
	})
})
