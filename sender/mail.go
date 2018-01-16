package sender

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"time"

	"fmt"
)

type Mail struct {
	From        string   `json:"from"`
	To          []string `json:"to"`
	Object      string   `json:"object"`
	Content     string   `json:"content"`
	Hash        string   `json:"hash,omitempty"`
	Timestamp   string   `json:"timestamp"`
	Environment string   `json:"environment,omitempty"`
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
		Hash:      fmt.Sprintf("%s", sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%s", from, object, content)))),
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
	}
}

func NewMails(data []byte) ([]Mail, error) {
	/*
		Create a new array from a given json
		parameter: <[]byte> Json array
		return: <[]Mail> A new array of mail
	*/
	mails := make([]Mail, 0)
	err := json.Unmarshal(data, &mails)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return mails, nil
}
