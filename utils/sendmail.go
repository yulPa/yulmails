package main

import (
	"log"
	"net/smtp"
)

func main() {
	to := []string{"recipient@arch-x250"}
	msg := []byte("To: recipient@arch-x250\r\n" +
		"Subject: test send mail\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	if err := smtp.SendMail("127.0.0.1:2525", nil, "sender@example.org", to, msg); err != nil {
		log.Fatal(err)
	}
}
