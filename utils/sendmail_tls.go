package main

import (
    "fmt"
    "log"
    "net"
    "net/mail"
    "net/smtp"
    "crypto/tls"
)

// StartTLS Email Example

func main() {

    from := mail.Address{"TestingSender", "username@example.tld"}
    to   := mail.Address{"TestingReceiver", "username@anotherexample.tld"}
    subj := "This is the email subject"
    body := "This is an example body.\n With two lines."

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    servername := "127.0.0.1:12800"

    host, _, _ := net.SplitHostPort(servername)

    auth := smtp.PlainAuth("","username@domain.com", "password", host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }

    c, err := smtp.Dial(servername)
    if err != nil {
        log.Panic(err)
    }

    c.StartTLS(tlsconfig)

    // Auth
    if err = c.Auth(auth); err != nil {
        log.Panic(err)
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
    }

    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
    }

    err = w.Close()
    if err != nil {
        log.Panic(err)
    }

    c.Quit()

}
