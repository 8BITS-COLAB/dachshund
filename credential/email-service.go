package credential

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	client *gomail.Dialer
}

func NewEmailService() *EmailService {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	return &EmailService{
		client: gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASS")),
	}
}

func (e *EmailService) SendWelcome(c *Credential) error {
	d := SendEmailDTO{
		Context: Context{
			To: c,
		},
	}

	t, _ := template.ParseFiles("asset/template/welcome.html")

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, d); err != nil {
		return err
	}

	result := tpl.String()
	msg := gomail.NewMessage()
	msg.SetHeader("Subject", "Welcome to Dachshund")
	msg.SetHeader("From", "Dachshund <no-reply@dachshund.io>")
	msg.SetHeader("To", fmt.Sprintf("%s <%s>", c.Name, c.Email))
	msg.SetBody("text/html", result)

	return e.client.DialAndSend(msg)
}

func (e *EmailService) SendDefault(c *Credential, d *SendEmailDTO) error {
	t, _ := template.ParseFiles("asset/template/default.html")

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, d); err != nil {
		return err
	}

	result := tpl.String()
	msg := gomail.NewMessage()
	msg.SetHeader("Subject", d.Context.Body.Subject)
	msg.SetHeader("From", fmt.Sprintf("%s <%s>", c.Name, c.Email))
	msg.SetHeader("To", fmt.Sprintf("%s <%s>", d.Context.To.Name, d.Context.To.Email))
	msg.SetBody("text/html", result)

	return e.client.DialAndSend(msg)
}
