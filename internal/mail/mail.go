package mail

import (
	"crypto/tls"

	"github.com/Devil666face/avzserver/internal/config"
	"github.com/gofiber/fiber/v2/log"
	gomail "gopkg.in/mail.v2"
)

type Smtp struct {
	from   string
	dialer *gomail.Dialer
}

// _username=user
// _password=123456
// _reciver=smtp@local.lan
// _from=user@local.lan
// _port=587
func New(_config *config.Config) *Smtp {
	d := gomail.NewDialer(
		_config.SMTPReciver,
		_config.SMTPPort,
		_config.SMTPUser,
		_config.SMTPPassword,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Smtp{
		dialer: d,
		from:   _config.SMTPEmail,
	}
}

func (s *Smtp) Send(to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Gomail test subject")
	m.SetBody("text/plain", "This is Gomail test body")
	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (s *Smtp) MustSend(to string) {
	if err := s.Send(to); err != nil {
		log.Info(err)
	}
}
