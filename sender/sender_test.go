package sender

import (
	"github.com/stretchr/testify/assert"

	"net/smtp"
	"testing"
)

type emailRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	body []byte
}

func mockSend(err error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, body []byte) error {
		*r = emailRecorder{addr, a, from, to, body}
		return err
	}, r
}

func TestSendMail(t *testing.T) {
	f, r := mockSend(nil)
	sender := &emailSender{
		conf: EmailConfig{
			ServerHost: "smtp.coucou",
			ServerPort: "123",
			Username:   "user",
			Password:   "pass",
			SenderAddr: "sender@address.com",
		},
		send: f,
	}
	err := sender.Send(Mail{
		From: "sender@address.com",
		To: []string{
			"sender@address.com",
			"sender1@address.com",
		},
		Object:  "An object",
		Content: "A content",
	})
	assert.Equal(t, r.from, "sender@address.com")
	assert.Nil(t, err)
}
