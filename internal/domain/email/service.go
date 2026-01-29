package email

import (
	"bytes"
	"html/template"
	"time"

	"eclaim-workshop-deck-api/internal/config"

	gomail "gopkg.in/gomail.v2"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendResetPassword(toEmail, username, newPassword string) error {
	tmpl, err := template.ParseFiles("templates/reset_password.html")
	if err != nil {
		return err
	}

	data := struct {
		Username string
		Password string
		Year     int
	}{
		Username: username,
		Password: newPassword,
		Year:     time.Now().Year(),
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// 2. Create the message
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailData.SMTP.Name+" <"+config.EmailData.SMTP.User+">")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your New Password")
	m.SetBody("text/html", body.String())

	// 3. Set up the dialer
	d := gomail.NewDialer(
		config.EmailData.SMTP.Server,
		int(config.EmailData.SMTPPort),
		config.EmailData.SMTP.User,
		config.EmailData.SMTP.Pass,
	)

	// 4. Send
	return d.DialAndSend(m)
}
