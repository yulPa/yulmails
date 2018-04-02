package spam

import (
  "os/exec"
  "bytes"
  "time"

  "github.com/robfig/cron"

  "github.com/yulPa/yulmails/pkg/logger"
)

var log = logger.GetLogger("ym-compute")

func checkSpamAssassin()  {
  /*
    This init method is called each time this package is loaded
    We check if `spamassassin` is installed on the machine
  */

}

func computeMail()  {
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
		time.Sleep(1000*time.Millisecond)
	}
}
