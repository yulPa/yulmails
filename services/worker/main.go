package worker

import (
	"github.com/adjust/rmq"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Configuration struct {
	QueueAddr      string `json:"queue_addr"`
	QueueName      string `json:"queue_name"`
	StoreQueueName string `json:"store_queue_name"`
	NbConsumer     int    `json:"nb_consumer"`
}

type Consumer struct {
	name    string
	count   int
	before  time.Time
	conf    *Configuration
	plugins []*plugin
}

func NewConsumer(tag int, c *Configuration, p []*plugin) *Consumer {
	return &Consumer{
		name:    fmt.Sprintf("worker-%d", tag),
		count:   0,
		before:  time.Now(),
		conf:    c,
		plugins: p,
	}
}

func NewConfiguration(path string) (*Configuration, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Configuration
	if err = json.Unmarshal(f, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Consumer) Consume(del rmq.Delivery) {
	c.count++
	// start the logic to perform YM checks
	m := del.Payload()
	res, err := c.plugins[0].SendEmail(m)
	if err != nil {
		log.Printf("consumer: %s - unable to check email: %v", c.name, err)
	}
	fmt.Print(string(res))
	del.Ack()
	// now we can store the email if its ready to be sent
	if ok := c.storeEmail(m); !ok {
		log.Printf("consumer: %s - unable to store email", c.name)
	}
}

func (c *Consumer) storeEmail(m string) bool {
	connection := rmq.OpenConnection("emailsService", "tcp", c.conf.QueueAddr, 1)
	queue := connection.OpenQueue(c.conf.StoreQueueName)
	return queue.Publish(m)
}

// StartWorker will start a worker in order to consume DB with consumers
func StartWorker(confPath string) error {
	conf, err := NewConfiguration(confPath)
	if err != nil {
		return err
	}
	// We should create the pipeline here
	plugins := []*plugin{NewPlugin("spamassassin-plugin:12800")}
	connection := rmq.OpenConnection("emailsService", "tcp", conf.QueueAddr, 1)
	queue := connection.OpenQueue(conf.QueueName)
	queue.StartConsuming(1000, 500*time.Millisecond)
	host, _ := os.Hostname()
	for i := 0; i < conf.NbConsumer; i++ {
		log.Printf("adding consumer: %d to worker: %s\n", i, host)
		queue.AddConsumer(host, NewConsumer(i, conf, plugins))
	}
	select {}
}
