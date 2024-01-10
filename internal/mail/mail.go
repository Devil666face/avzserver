package mail

import (
	"crypto/tls"

	"github.com/Devil666face/avzserver/internal/config"
	"github.com/gofiber/fiber/v2/log"
	gomail "gopkg.in/mail.v2"
)

type Mail struct {
	from   string
	dialer *gomail.Dialer
}

// _username=user
// _password=123456
// _reciver=smtp@local.lan
// _from=user@local.lan
// _port=587
func New(_config *config.Config) *Mail {
	d := gomail.NewDialer(
		_config.SMTPReciver,
		_config.SMTPPort,
		_config.SMTPUser,
		_config.SMTPPassword,
	)
	//nolint:gosec // Self-signed certs
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Mail{
		dialer: d,
		from:   _config.SMTPEmail,
	}
}

func (s *Mail) Send(to, href string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Активация аккаунта")
	m.SetBody("text/plain", href)
	return s.dialer.DialAndSend(m)
}

func (s *Mail) MustSend(to, href string) {
	if err := s.Send(to, href); err != nil {
		log.Info(err)
	}
}
