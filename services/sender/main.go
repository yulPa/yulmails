package sender

import (
	"github.com/adjust/rmq"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strings"
	"time"
)

type Configuration struct {
	QueueAddr  string `json:"queue_addr"`
	QueueName  string `json:"queue_name"`
	NbConsumer int    `json:"nb_consumer"`
}

type Consumer struct {
	name   string
	count  int
	before time.Time
}

func NewConsumer(tag int) *Consumer {
	return &Consumer{
		name:   fmt.Sprintf("worker-%d", tag),
		count:  0,
		before: time.Now(),
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
	fmt.Print(del.Payload())
	del.Ack()
}

func listRecipients(msg string) ([]string, error) {
	m, err := mail.ReadMessage(strings.NewReader(msg))
	if err != nil {
		return nil, err
	}
	header := m.Header
	addr, err := header.AddressList("To")
	if err != nil {
		return nil, err
	}
	recipients := make([]string, 0, len(addr))
	for _, a := range addr {
		recipients = append(recipients, a.Address)
	}
	return recipients, nil
}

// StartSender will start a sender in order to consume DB with consumers and send them
func StartSender(confPath string) error {
	conf, err := NewConfiguration(confPath)
	if err != nil {
		return err
	}
	connection := rmq.OpenConnection("emailsService", "tcp", conf.QueueAddr, 1)
	queue := connection.OpenQueue(conf.QueueName)
	queue.StartConsuming(1000, 500*time.Millisecond)
	host, _ := os.Hostname()
	for i := 0; i < conf.NbConsumer; i++ {
		log.Printf("adding consumer: %d to sender: %s\n", i, host)
		queue.AddConsumer(host, NewConsumer(i))
	}
	select {}
}
