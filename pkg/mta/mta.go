package mta

import (
	"log"
	"time"

	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/yulPa/yulmails/pkg/mta/processor"
)

var (
	cfg = &guerrilla.AppConfig{
		LogFile:      "/tmp/guerrilla.log",
		AllowedHosts: []string{"."},
		BackendConfig: backends.BackendConfig{
			"save_process": "Hasher|RedisQueue",
			"redis_addr":   "redis:6379",
		},
		Servers: []guerrilla.ServerConfig{guerrilla.ServerConfig{
			ListenInterface: "0.0.0.0:25",
			IsEnabled:       true,
		}},
	}
	d = guerrilla.Daemon{Config: cfg}
)

// Run will start the server
func Run() {
	d.AddProcessor("RedisQueue", processor.RedisQueueProcessor)

	if err := d.Start(); err != nil {
		log.Fatalf("unable to start mta: %v", err)
	}
	for {
		time.Sleep(60 * time.Minute)
	}

}
