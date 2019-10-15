package main

import (
	"log"
	"net/smtp"
)

func main() {
	to := []string{"recipient@arch-x250"}
	msg := []byte(`From: bkd@limapuluhkotakab.go.id
Content-Transfer-Encoding: base64
Content-Type: text/plain; charset=UTF-8
Mime-Version: 1.0 
Subject: Margaretta
Message-Id: <1805078D-22A4-F9AA-FD80-089AFFFBBFD9@limapuluhkotakab.go.id>
Date: Tue, 29 May 2018 21:14:02 +0200
To: recipient@arch-x250


VGVsbCBtZSwgdGVtcHRlciwgd2hhdCBkbyB5b3UgaGF2ZSBpbiB5b3VyIHBhbnRzPw0KSXMgaXQg
YSBoYXJkIGluc3RydW1lbnQgdGhhdCBjYW4gZ2l2ZSBtZSBwbGVhc3VyZT8NCkhvcGUsIGl0IGlz
LiBDb21lIGFuZCBwcm92ZSB0aGF0IHlvdSBhcmUgYSByZWFsIHN0YWxsaW9uIQ0KDQpodHRwOi8v
d3d3LmZpbmRjb21wdXRlcmdlZWsuY29tL2J2Z3k3OXgvd3NmZnMwY3oucGhwP1kyOXVkR0ZqZEVC
NWRXeHdZUzVwYnc9PQ0K
`)
	if err := smtp.SendMail("127.0.0.1:2525", nil, "sender@example.org", to, msg); err != nil {
		log.Fatal(err)
	}
}
