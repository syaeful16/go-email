package main

import (
	"log"

	"github.com/yourusername/go-email"
)

func main() {
	smtpClient := email.NewSMTPClient("smtp.example.com", 587, "user", "pass", "from@example.com")

	err := smtpClient.Send("to@example.com", "Hello", "This is a test email")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent!")
}
