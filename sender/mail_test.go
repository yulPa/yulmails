package sender

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
)

func TestCreateAMail(t *testing.T) {
	m := NewMail(
		"sender@mail.com",
		[]string{
			"recipipents1@mail.com",
			"recipipents2@mail.com",
		},
		"A new email",
		"This is a content",
	)
	assert.Len(t, m.To, 2)
	assert.Equal(t, fmt.Sprintf("%x", m.Hash), "c009181ee42b2b5e8e11fbe9883da12861845cc874baf6c15ee6527bcc45af5d")
}
