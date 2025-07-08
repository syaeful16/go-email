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

func (s *SMTPClient) Send(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	msg := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n",
	)
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	return smtp.SendMail(addr, auth, s.From, []string{to}, msg)
}
