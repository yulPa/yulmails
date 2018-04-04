package mail

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"testing"
	"net/mail"
	"strings"
)

func TestCreateAMail(t *testing.T) {
	raw := `Return-Path: email@yulmails.io
To: emailTo@yulmails.io
From: emailFrom <emailFrom@yulmails.io>
Subject: coucou
Openpgp: preference=signencrypt
Autocrypt: addr=mathieu.tortuyaux2@etu.univ-lorraine.fr; keydata=
 xsBNBFqartoBCADE/P4lTWv/gvFZwWVxh0KgLTLbkS3xNJBumkTN1cndvEXx0f544JKSAVML
 8XMT560OHIqLtAH+sz2ldLpWjMJ06xaX0AyisBbig2aCWqTPUBRT8/ujK/dbULsl5gk1bGC9
 7oCfKOtiJMwP3aa3hSwqVLuwz/lqZX8+lZe9g1EXbMa7XFDAT6qtcfsRK5opf3yo2zfY7IEq
 ATGsEGbXl5hCV+rNR1V9eY5PFRvbQp6+uoJJda2+0DsQeL1ORF05daQLKF1f7Oslnth2munu
 vycDGY2IT0PKkU4SKTeVKAXbA9AeCn2ksYxwWFBkWY364ZaGFa/CsVpE7GzXvgIuUnyBABEB
 AAHNO01hdGhpZXUgVG9ydHV5YXV4IDxtYXRoaWV1LnRvcnR1eWF1eDJAZXR1LnVuaXYtbG9y
 cmFpbmUuZnI+wsCUBBMBCAA+FiEEyGP8e/7ZmmJmbESrNoJk4U9/MuoFAlqartoCGwMFCQPC
 ZwAFCwkIBwIGFQoJCAsCBBYCAwwdECHgECF4AACgkQNoJk4U9/MupabggAo4SX2w4DN/K2DT
 qsYjQaBii4gms3GaIXDGpRSHbNMLeXW0MXtGPbUleM9Pw6ilyqypIfvD6jVExUoO906mPyKZ
 MCFIrAC+YAOriIA52yWPVyrHIZ0eCJNUXfbrI2GbF6c4qd1XTTmcM2tSOLhdpJn8LVaWD0W8
 2y3lrs3nXZD90nZITfZMqYNWUG5spbkO9USfFmgOKAGRRMlkoeIaYQrif6YddV5YOoYI9+ql
 7Lkg3Pm9mFEnVic1nusqM5F83d/uASHdWI+bbW49MgxTs3Cd4qYYdO/o+Z04Lp/ISYAioRBC
 qg9ggA7THTv/LuWv28ZgfBHBQd3AIZAoUtEZ/87ATQRamq7aAQgA33i66ev4fBe4EuxpRSx7
 DIW1eqAkRJIDRFgrECseddPfgXl901neemHTKYZsdjFfCcTeJOgOM7piqVx35A2ysAHaIk8A
 F7Gy4dEo69AcbBOYR0q/1TmQ+dpmxQeWCFFibEoFC+ZwGV8nyhV+yB+7RLF6Q0Cqnhtjd0He
 j6P3tOtNfhoC+vuNg/X4xzQRDWZPi9tJoLlnOp/U3+Ge/ysRunC65ByWIBoCISUcoVQTZT5m
 MVgiQUXp3XveHBv8XtFod5TPyoglju5RCnhRST2ZMV/Ch3if8CagMO4n0ZBUFiY+YsXLqpPK
 YUYAga2YAF0p5JeSp8oDxq/Rv2YvZt6H6QARAQABwsB8BBgBCAAmFiEEyGP8e/7ZmmJmbESr
 NoJk4U9/MuoFAlqartoCGwwFCQPCZwAACgkQNoJk4U9/MupOqQf/c+vPBakz2yO3+U5tadjm
 I9hLj/OFTrsNpUMSGNSt61o79lU0Lp+uIJt/xWxdwJGjCW4fyJAITFzxLXayDSfhDQaS/aRY
 ibwqT7lO0bZsReAAgf/xnvqn4Q4moNubaiwC8VZjsayuN/7fxv3jxFhHi28sSGk1SovKqHEo
 34iXPJKGgKqNeRyardmLFOvRMpS503zhV5x7qwOnSTuSIrhPFTlt8Kr2kSZFPE8lNocYcNHr
 INK1ame86IlrEhw3yb/eGCi+dj0pnBPSkAcXE9aCkCg0cg1NW7ibQGJ7fErAKwdqcIvD1VoR
 nNA4NW+GwcjbMQ9foIn4KZga5RMUq5Jizg==
Message-ID: <eb07f120-7371-01ce-4568-5d51da0e587c@etu.univ-lorraine.fr>
Date: Tue, 27 Mar 2018 21:57:42 +0200
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101
 Thunderbird/52.7.0
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 7bit
Content-Language: fr-LU

slt sa va
`
	r := strings.NewReader(raw)
	m, _ := mail.ReadMessage(r)
	mEntry := NewMailEntry(m)
	assert.Equal(t, mEntry.Message.Header.Get("To"), "emailTo@yulmails.io")
	assert.Equal(t, fmt.Sprintf("%x", mEntry.Hash), "95e1c2f0155b76fd0b7cec4f77d8abc58b311af9ed59eafe85bf9074024bcd76")
}
