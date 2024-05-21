package main

import (
	"log"
	"net/smtp"
)

func main() {

	// Mailtrap account config

	username := "api"

	password := "a0497b5b3fe0f0f5d4b878bd9f33532a"

	smtpHost := "live.smtp.mailtrap.io"

	// Choose auth method and set it up

	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Message data

	from := "mailtrap@demomailtrap.com"

	to := []string{"carl.klagba@gmail.com"}

	message := []byte("To: carl.klagba@gmail.com\r\n" +
		"From: mailtrap@demomailtrap.com\r\n" +
		"Subject: Why aren't you using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here's the space for your great sales pitch\r\n")

	// Connect to the server and send message

	smtpUrl := smtpHost + ":587"

	err := smtp.SendMail(smtpUrl, auth, from, to, message)

	if err != nil {

		log.Fatal(err)

	}
}
