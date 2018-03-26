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

func TestcreateMails(t *testing.T) {
	data := []byte(`
		[
		  {
		      "from": "sender@mail.com",
		      "to": [
		        "recipipents1@mail.com",
		        "recipipents2@mail.com"
		      ],
		      "object": "A new email",
		      "content": "This is a content",
		      "hash": "c009181ee42b2b5e8e11fbe9883da12861845cc874baf6c15ee6527bcc45af5d",
		      "timestamp": "1516017704"
		  },
		  {
		      "from": "another@mail.com",
		      "to": [
		        "recipipents1@mail.com"
		      ],
		      "object": "Another new email",
		      "content": "This is a content",
		      "hash": "5c6a9c8af487bb937bd06e60236e5e63ffb04919ee5e157832f87b6e9220a2c5",
		      "timestamp": "1516017764"
		  }
		]
		`)
	m, e := NewMails(data)
	assert.Len(t, m, 2)
	assert.Nil(t, e)
}
