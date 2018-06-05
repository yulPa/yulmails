package main

import (
	"log"
	"time"

	"github.com/flashmob/go-guerrilla/backends"

	"github.com/flashmob/go-guerrilla"

	"github.com/yulPa/yulmails/pkg/mta/processor"
)

var (
	cfg = &guerrilla.AppConfig{
		LogFile:      "/tmp/guerrilla.log",
		AllowedHosts: []string{"arch.localdomain"},
		BackendConfig: backends.BackendConfig{
			"save_process": "YulmailsProcessor",
		},
	}
	d = guerrilla.Daemon{Config: cfg}
)

func main() {
	d.AddProcessor("YulmailsProcessor", processor.YulmailsProcessor)
	if err := d.Start(); err != nil {
		log.Fatalf("unable to start mta: %v", err)
	}

	for {
		time.Sleep(60 * time.Minute)
	}

}
