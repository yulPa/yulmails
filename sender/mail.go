package sender

import (
	"crypto/sha256"
	"time"

	"fmt"
)

type Mail struct {
	From        string            `json:"from"`
	To          []string          `json:"to"`
	Object      string            `json:"object"`
	Content     string            `json:"content"`
	Hash        [sha256.Size]byte `json:"hash"`
	Timestamp   int64             `json:"timestamp"`
	Environment string            `json:"environment,omitempty"`
}

func NewMail(from string, to []string, object string, content string) *Mail {
	/*
	   Create a new mail
	   parameter: <string> Sender
	   parameter: <[]string> recipipents
	   parameter: <string> Object
	   parameter: <string> content
	*/
	return &Mail{
		From:      from,
		To:        to,
		Object:    object,
		Content:   content,
		Hash:      sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%s", from, object, content))),
		Timestamp: time.Now().Unix(),
	}
}
