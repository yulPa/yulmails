package main

import (
	"testing"
	"net/mail"
	"strings"

	"github.com/stretchr/testify/assert"
)

func TestGetScore(t *testing.T){
	m := `From: "TestingSender" <username@example.tld>
To: "TestingReceiver" <username@anotherexample.tld>
Subject: This is the email subject
X-Spam-Status:Yes, score=7.9 required=5.0 tests=EMPTY_MESSAGE,MISSING_DATE, MISSING_FROM,MISSING_HEADERS,MISSING_MID,MISSING_SUBJECT, NO_HEADERS_MESSAGE,NO_RECEIVED,NO_RELAYS autolearn=no autolearn_force=no version=3.4.2
X-Spam-Checker-Version:SpamAssassin 3.4.2 (2018-09-13) on spamassassin-plugin
X-Spam-Level:*******

This is an example body.
With two lines.
`
	r, _ := mail.ReadMessage(strings.NewReader(m))
	s, _:= getScore(r)
	assert.Equal(t, 7.9, s)
}
