package notification

import (
	"fmt"
	"github.com/CarlKlagba/go-todo/repository"
	"log"
	"net/smtp"
	"strings"
)

func SendNotification(tasks []repository.Task) {
	var strBuilder strings.Builder
	for _, task := range tasks {
		strg := fmt.Sprintf("Task %s is due in %d days \n", task.Task, task.DaysLeft())
		strBuilder.WriteString(strg)
	}

	mailBody := strBuilder.String()

	/// Mailtrap account config
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
		"Subject: Your tasks due soon !\r\n" +
		"\r\n" +
		mailBody)

	// Connect to the server and send message

	smtpUrl := smtpHost + ":587"

	log.Println(tasks)
	log.Println("MailBody " + mailBody)
	log.Println("message: " + string(message))

	err := smtp.SendMail(smtpUrl, auth, from, to, message)

	if err != nil {
		log.Println("Error: " + err.Error())
		log.Fatal(err)

	}
}
