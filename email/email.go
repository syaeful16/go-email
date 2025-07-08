package email

// EmailSender defines interface untuk pengiriman email
type EmailSender interface {
	Send(to string, subject string, body string) error
}
