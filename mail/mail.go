package mail

import (
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

// Config ...
type Config struct {
	SMTPAddr string
	Identity string
	UserName string
	Password string
	Host     string
}

// Send ...
func Send(config Config, from, to, subject, text, html string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	// e.Subject = `subject`
	e.Text = []byte(text)
	// e.HTML = []byte(html)
	err := e.Send(
		config.SMTPAddr,
		smtp.PlainAuth(config.Identity, config.UserName, config.Password, config.Host),
		// `smtp.163.com:25`,
		// smtp.PlainAuth(``, `imbotbotbot@163.com`, `VNAKFMVXQOJGQNGK`, `smtp.163.com`),
	)
	if err != nil {
		logrus.WithError(err).Error()
	}
	return err
}
