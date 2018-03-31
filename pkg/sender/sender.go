package sender

import (
	"github.com/yulPa/yulmails/pkg/logger"

	"github.com/robfig/cron"

	"fmt"
	"net/smtp"
	"time"
)

var log = logger.GetLogger("ym-sender")

type EmailSender interface {
	Send(Mail) error
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

func (e *emailSender) Send(mail Mail) error {
	/*
		Send an email to a given list of recipipents
		parameter: <[]string> Array of recipipents
		parameter: <[]byte> Body content
		return: <error> Return nil if no errors
	*/
	addr := fmt.Sprintf("%s:%s", e.conf.ServerHost, e.conf.ServerPort)
	auth := smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost)
	return e.send(addr, auth, e.conf.SenderAddr, mail.To, []byte(fmt.Sprintf("%s%s", mail.Object, mail.Content)))
}

func sendMail() {
	/*
		Request mails that are ready to be sent
	*/
	/*
	var cli = client.NewHTTPClient()

	req, _ := http.NewRequest(http.MethodGet, "yulmails-api:9252/v1/mails", nil)
	res, err := cli.Do(req)

	if err != nil {
		log.Infoln(err)
	}
	log.Infoln(res)
	
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()


	
	mails, err := NewMails(body)

	for _, m := range mails {
		fmt.Println("Send mail", m)
	}
	*/
	for i := 0; i < 5; i++ {
		fmt.Println("Sending mail")
	}
}

func Run() {
	var c = cron.New()

	c.AddFunc("@every 5s", sendMail)

	log.Infoln("Sender Cron has started")
	c.Start()

	for {
		time.Sleep(1000*time.Millisecond)
	}
}
