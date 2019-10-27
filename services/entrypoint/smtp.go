package entrypoint

import (
	"errors"
	"fmt"
	"time"

	"github.com/flashmob/go-guerrilla"
	"github.com/yulpa/yulmails/services/entrypoint/processor"
)

// StartSMTP will start the SMTP server with the configuration file
func StartSMTP(confPath string) error {
	d := guerrilla.Daemon{}
	if _, err := d.LoadConfig(confPath); err != nil {
		return errors.New(fmt.Sprintf("unable to load SMTP config: %v", err))
	}
	d.AddProcessor("RedisQueue", processor.RedisQueueProcessor)

	if err := d.Start(); err != nil {
		return errors.New(fmt.Sprintf("unable to start SMTP: %v", err))
	}
	for {
		time.Sleep(60 * time.Minute)
	}
}
