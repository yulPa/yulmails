package sender

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRecipients(t *testing.T) {

	tests := []struct {
		mail       string
		recipients []string
	}{
		{
			mail: `From: bkd@limapuluhkotakab.go.id
Content-Transfer-Encoding: base64
Content-Type: text/plain; charset=UTF-8
Mime-Version: 1.0 
Subject: Margaretta
Message-Id: <1805078D-22A4-F9AA-FD80-089AFFFBBFD9@limapuluhkotakab.go.id>
Date: Tue, 29 May 2018 21:14:02 +0200
To: recipient@arch-x250


VGVsbCBtZSwgdGVtcHRlciwgd2hhdCBkbyB5b3UgaGF2ZSBpbiB5b3VyIHBhbnRzPw0KSXMgaXQg
`,
			recipients: []string{"recipient@arch-x250"},
		},
		{
			mail: `From: bkd@limapuluhkotakab.go.id
Content-Transfer-Encoding: base64
Content-Type: text/plain; charset=UTF-8
Mime-Version: 1.0 
Subject: Margaretta
Message-Id: <1805078D-22A4-F9AA-FD80-089AFFFBBFD9@limapuluhkotakab.go.id>
Date: Tue, 29 May 2018 21:14:02 +0200
To: recipient@arch-x250, dest@local.tld, toto@superman.io


VGVsbCBtZSwgdGVtcHRlciwgd2hhdCBkbyB5b3UgaGF2ZSBpbiB5b3VyIHBhbnRzPw0KSXMgaXQg
`,
			recipients: []string{"recipient@arch-x250", "dest@local.tld", "toto@superman.io"},
		},
		{
			mail: `From: bkd@limapuluhkotakab.go.id
Content-Transfer-Encoding: base64
Content-Type: text/plain; charset=UTF-8
Mime-Version: 1.0 
Subject: Margaretta
Message-Id: <1805078D-22A4-F9AA-FD80-089AFFFBBFD9@limapuluhkotakab.go.id>
Date: Tue, 29 May 2018 21:14:02 +0200
To:


VGVsbCBtZSwgdGVtcHRlciwgd2hhdCBkbyB5b3UgaGF2ZSBpbiB5b3VyIHBhbnRzPw0KSXMgaXQg
`,
			recipients: []string{},
		},
	}
	for _, test := range tests {
		recipients, _ := listRecipients(test.mail)
		assert.Equal(t, len(recipients), len(test.recipients))
	}
}
