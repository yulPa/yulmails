package processing

import (
	"fmt"
	"log"
	"net/mail"
	"os/exec"
	"strings"
	"time"

	"github.com/adjust/rmq"
)

var (
	connection rmq.Connection
)

type consumer struct {
	name   string
	count  int
	before time.Time
}

func newConsumer(tag int) *consumer {
	return &consumer{
		name:   fmt.Sprintf("consumer-%d", tag),
		count:  0,
		before: time.Now(),
	}
}

func (cons *consumer) Consume(delivery rmq.Delivery) {
	cons.count++
	if cons.count%10 == 0 {
		cons.before = time.Now()
		log.Println(delivery.Payload())
	}
	time.Sleep(time.Millisecond)
	if cons.count%10 == 0 {
		delivery.Reject()
	} else {
		delivery.Ack()
	}
}

func init() {
	connection = rmq.OpenConnection("emailsService", "tcp", "redis:6379", 1)
}

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

// Run will start our consumer
func Run() {
	queue := connection.OpenQueue("emails")
	queue.StartConsuming(100, 500*time.Millisecond)
	queue.AddConsumer("batch-consumer", newConsumer(1))
	select {}
}
