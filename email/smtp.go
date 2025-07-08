package email

import (
	"fmt"
	"net/smtp"
)

type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewSMTPClient(host string, port int, username, password, from string) *SMTPClient {
	return &SMTPClient{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *SMTPClient) Send(to, subject, htmlBody string) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	// Tambahkan header email HTML
	msg := []byte(
		"From: " + s.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			htmlBody + "\r\n",
	)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	return smtp.SendMail(addr, auth, s.From, []string{to}, msg)
}
