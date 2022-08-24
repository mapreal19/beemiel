package mailer

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/mailgun/mailgun-go"
)

var Region = struct {
	EU, US string
}{
	EU: "EU",
	US: "US",
}

type mailGun struct {
	globalConf
	Region, Domain string
}

var mailGunConf mailGun

func SetMailgun(region, domain string) {
	mailGunConf.Domain = domain
	mailGunConf.Region = region
}

func newMailgunSender(g globalConf) *emailSender {
	mailGunConf.globalConf = g
	return &emailSender{mailGunConf.mailgunSender}
}

func (m *mailGun) mailgunSender(email Email) error {
	beego.Info("Sending email through MailGun... Recipient: ", email.Tos[0])
	mg := mailgun.NewMailgun(m.Domain, m.ApiKey)
	m.setRegion(mg, m.Region)

	message := mg.NewMessage(
		fmt.Sprintf("%s<%s>", email.FromName, email.From),
		email.Subject,
		email.PlainBody,
		email.Tos...)

	m.addCCs(message, email.Ccs)
	m.addBCCs(message, email.Bccs)
	m.addAttachments(message, email.Attachments)
	message.SetHtml(email.Body)

	_, _, err := mg.Send(message)
	return err
}

func (m *mailGun) setRegion(mg *mailgun.MailgunImpl, region string) {
	switch region {
	case Region.EU:
		mg.SetAPIBase("https://api.eu.mailgun.net/v3")
	case Region.US:
		mg.SetAPIBase("https://api.mailgun.net/v3")
	}
}

func (m *mailGun) addAttachments(message *mailgun.Message, attachment []Attachment) {
	for _, att := range attachment {
		message.AddBufferAttachment(att.Name, att.Data)
	}
}

func (m *mailGun) addCCs(message *mailgun.Message, addresses []string) {
	for _, addr := range addresses {
		message.AddCC(addr)
	}
}

func (m *mailGun) addBCCs(message *mailgun.Message, addresses []string) {
	for _, addr := range addresses {
		message.AddBCC(addr)
	}
}
