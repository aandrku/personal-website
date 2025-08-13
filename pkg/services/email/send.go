package email

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aandrku/personal-website/pkg/view/shared"
	"gopkg.in/gomail.v2"
)

const (
	host = "smtp.gmail.com"
	port = 587
)

func SendContact(name string, email string, message string) error {
	emailSender := os.Getenv("EMAIL_SENDER")
	emailRecepient := os.Getenv("EMAIL_RECEPIENT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	d := gomail.NewDialer(host, port, emailSender, emailPassword)

	body, err := renderContactEmail(name, email, message)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailRecepient)
	m.SetHeader("Subject", fmt.Sprintf("%s is trying to reach me.", name))
	m.SetBody("text/html", body)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func renderContactEmail(name, email, message string) (string, error) {
	var buf bytes.Buffer
	err := shared.Email(name, email, message).Render(context.Background(), &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendString(subject, str string) error {
	emailSender := os.Getenv("EMAIL_SENDER")
	emailRecepient := os.Getenv("EMAIL_RECEPIENT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")

	d := gomail.NewDialer(host, port, emailSender, emailPassword)

	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailRecepient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", str)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil

}
