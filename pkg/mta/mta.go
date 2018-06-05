package main

import (
	"log"
	"time"

	"github.com/flashmob/go-guerrilla"
)

var (
	cfg = &guerrilla.AppConfig{
		LogFile:      "/tmp/guerrilla.log",
		AllowedHosts: []string{"arch.localdomain"},
	}
	d = guerrilla.Daemon{Config: cfg}
)

func main() {
	if err := d.Start(); err != nil {
		log.Fatalf("unable to start mta: %v", err)
	}

	for {
		time.Sleep(60 * time.Minute)
	}

}
