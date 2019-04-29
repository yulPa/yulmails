package processing

import (
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/adjust/rmq"
)

var (
	connection rmq.Connection
)

type batchConsumer struct {
	name   string
	count  int
	before time.Time
}

func newBatchConsumer(tag int) *batchConsumer {
	return &batchConsumer{
		name:   fmt.Sprintf("consumer-%d", tag),
		count:  0,
		before: time.Now(),
	}
}

func (cons *batchConsumer) Consume(deliveries rmq.Deliveries) {
	for _, delivery := range deliveries {
		log.Println(delivery.Payload())
		delivery.Ack()
		checkSpamAssassin(delivery.Payload())
	}
	time.Sleep(time.Millisecond)
}

func init() {
	connection = rmq.OpenConnection("emailsService", "tcp", "redis:6379", 1)
}

func genUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func checkSpamAssassin(content string) {
	/*
	   This method will call SA in order to determine if mails are spams
	*/
	id := genUuid()
	f, err := os.Create(id)
	if err != nil {
		log.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		log.Println(err)
		f.Close()
		return
	}
	log.Println(l, "bytes written successfully to ", id)
	err = f.Close()
	if err != nil {
		log.Println(err)
		return
	}

	cmd := exec.Command("spamassassin", "-t", id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	} else {
		log.Printf("combined out:\n%s\n", string(out))

		raw := string(out)
		r := strings.NewReader(raw)
		m, _ := mail.ReadMessage(r)
		if strings.Split(m.Header.Get("X-Spam-Status"), ",")[0] == "Yes" {
			log.Println("Spam")
		}
		log.Println("not spam")
	}
	// delete file
	err = os.Remove(id)
	if err != nil {
		log.Printf("Error deleting file: ", err)
	} else {
		log.Println(id, "deleted")
	}
}

// Run will start our consumer
func Run() {
	queue := connection.OpenQueue("emails")
	queue.StartConsuming(100, 500*time.Millisecond)
	queue.AddBatchConsumer("batch-consumer", 10, newBatchConsumer(1))
	select {}
}
