package sender

import (
	"github.com/yulPa/yulmails/pkg/logger"
	"github.com/yulPa/yulmails/pkg/mongo"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"

	"fmt"
	"net/smtp"
	"strings"
	"time"
)

var log = logger.GetLogger("ym-sender")
var mailSender = NewMailSender(EmailConfig{
	ServerHost: "server.local.tld",
	ServerPort: "25",
	Username:   "username",
	Password:   "password",
	SenderAddr: "sender@local.tld",
})

type EmailSender interface {
	Send(MailEntry) error
}

type emailSender struct {
	conf EmailConfig
	send func(string, smtp.Auth, string, []string, []byte) error
}
type EmailConfig struct {
	ServerHost string
	ServerPort string
	Username   string
	Password   string
	SenderAddr string
}

func NewMailSender(conf EmailConfig) EmailSender {
	/*
		Create a new mail sender
		parameter: <EmailConfig> A config associated to the email sender
		return: <EmailSender> An email sender
	*/
	return &emailSender{conf, smtp.SendMail}
}

func (e *emailSender) Send(mailEntry MailEntry) error {
	/*
		Send an email to a given list of recipipents
		parameter: <[]string> Array of recipipents
		parameter: <[]byte> Body content
		return: <error> Return nil if no errors
	*/
	addr := fmt.Sprintf("%s:%s", e.conf.ServerHost, e.conf.ServerPort)
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, getListOfRecipients(mailEntry), []byte(fmt.Sprintf("%s%s", mailEntry.Message.Header.Get("Subject"), mailEntry.Message.Body)))
}

func getListOfRecipients(mailEntry MailEntry) []string {
	/*
		Return all recipipents for a given email
		parameter: <mailEntry> The given email
	*/
	return strings.Split(mailEntry.Message.Header.Get("To"), ",")
}

func sendMail() {
	/*
		Request mails that are ready to be sent
	*/
	var workdb = mongo.NewSession("mongodb://workdb:27017")
	dbMails := workdb.DB("buffer")

	mailsToSend := dbMails.GetSendableMails()

	for _, mail := range mailsToSend {
		if err = mailSender.Send(mail); err != nil {
			log.WithFields(logrus.Fields{
				"subject": mail.Message.Header.Get("subject"),
			}).Error("Error while sending mail. ", err)
		}
	}

}

func Run() {
	var c = cron.New()

	c.AddFunc("@every 5s", sendMail)

	log.Infoln("Sender Cron has started")
	c.Start()

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
