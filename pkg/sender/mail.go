package sender

import (
	"crypto/sha256"
	"strconv"
	"time"
	"fmt"
	"net/mail"
)

type MailEntry struct {
	*mail.Message
	Hash        string   `json:"hash,omitempty"`
	Timestamp   string   `json:"timestamp"`
	Environment string   `json:"environment,omitempty"`
}

func NewMailEntry(message *mail.Message) *MailEntry {
	/*
	   Create a new mail
		 parameter: <mail.Message> A simple message following this struct : Headers + Body
	*/
	return &MailEntry{
		Message: 	message,
		Hash:      fmt.Sprintf("%s", sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%s", message.Header.Get("From"), message.Header.Get("Object"), message.Body)))),
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
	}
}
