package mails

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAMessage(t *testing.T) {
	m := CreateAMessage(
		"sender@sender.com",
		[]string{"recipient1@recipient.com", "recipient2@recipient.com"},
		"A simple mail",
		"Wassup",
	)
	assert.NotNil(t, m)
}
