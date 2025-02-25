package utils

import (
	"log"

	"github.com/Amierza/warasin-mobile/backend/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(toEmail string, subject string, body string) error {
	emailConfig, err := config.NewEmailConfig()
	if err != nil {
		log.Printf("failed to load email config: %v", err)
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.AuthEmail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.AuthEmail,
		emailConfig.AuthPassword,
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Printf("failed to send email to %v: %v", toEmail, err)
		return err
	}

	return nil
}
