package processor

import (
	"log"

	"github.com/adjust/rmq"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
)

type redisQueueConfig struct {
	Addr string `json:"redis_addr"`
}

// RedisQueueProcessor is a custom processor to create a queue on redis
var RedisQueueProcessor = func() backends.Decorator {
	var config = &redisQueueConfig{}
	initFunc := backends.InitializeWith(func(backendConfig backends.BackendConfig) error {
		configType := backends.BaseConfig(&redisQueueConfig{})
		bcfg, err := backends.Svc.ExtractConfig(backendConfig, configType)
		if err != nil {
			return err
		}
		config = bcfg.(*redisQueueConfig)
		return nil
	})
	backends.Svc.AddInitializer(initFunc)
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskSaveMail {
					connection := rmq.OpenConnection("emailsService", "tcp", config.Addr, 1)
					queue := connection.OpenQueue("queue-entrypoint")
					if ok := queue.Publish(e.Data.String()); !ok {
						log.Println("unable to push email on queue")
					}
					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}
