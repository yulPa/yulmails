package mails

import "gopkg.in/gomail.v2"

func CreateAMessage(from string, to []string, subject string, body string) *gomail.Message {
	/*
	   This function will simply create a message following this structure
	   parameter: <string> Sender's address
	   parameter: <string> Recipient's address
	   parameter: <string> Mail's subject
	   parameter: <string> Mail's content
	   return: <gomail.Message> A gomail Message
	*/
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return m
}

// TODO: Add method in order to send mail with attachment
