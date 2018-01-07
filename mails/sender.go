package mails

import (
	"gopkg.in/gomail.v2"

	"github.com/check_mails/logger"
)

/* This file has not to be tested, `gomail` has his own test series */
var log = logger.GetLogger()

func SendMail(d *gomail.Dialer, m *gomail.Message) bool {
	/*
		Send a message with smtp server
		parameter: <gomail.Dialer> A struct to handler communication with smtp
		parameter: <gomail.Message> Message to send
		return: <bool> Return true if message has been sent
	*/
	if err := d.DialAndSend(m); err != nil {
		log.Errorln(err)
		return false
	}
	return true
}
