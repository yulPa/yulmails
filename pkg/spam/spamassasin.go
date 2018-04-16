package spam

import (
	"bytes"
	"net/mail"
	"os/exec"
	"strings"
	"time"

	"github.com/robfig/cron"

	"github.com/yulPa/yulmails/pkg/logger"
)

var log = logger.GetLogger("ym-compute")

func checkSpamAssassin(content string) {
	/*
	   This method will call SA in order to determine if mails are spams
	*/
	output, _ := exec.Command("spamassassin", content).Output()
	raw := string(output)

	r := strings.NewReader(raw)
	m, _ := mail.ReadMessage(r)
	if strings.Split(m.Header.Get("X-Spam-Status"), ",")[0] == "Yes" {
		log.Println("Spam")
	}
	log.Println("not spam")
}

func computeMail() {
	log.Infoln("Compute email")
}

func Run() {

	var out bytes.Buffer
	var c = cron.New()

	cmd := exec.Command("which", "spamassassin")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatal("`spamassassin` is not installed on this machine: ", err)
	}

	c.AddFunc("@every 5s", computeMail)

	log.Infoln("Compute Cron has started")
	c.Start()

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
