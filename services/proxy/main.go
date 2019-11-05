package proxy

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-smtp"
	"gopkg.in/yaml.v2"
)

// Configuration is the configuration for
// the mail proxy.
type Configuration struct {
	// Port of the proxy
	Port int `yaml:"port"`
	// TCP or UNIX address to listen on
	Addr string `yaml:"listen_address"`
	// MaxRecipients is the number of maximum recipients in the email
	MaxRecipients int `yaml:"max_recipients"`
	// YMAddr is the address of the YM entrypoint
	YMAddr string `yaml:"yulmails_address"`
	// TODO: add TLS configuration
}

type backend struct{ ymAddr string } //TODO: add API client here
type session struct {
	from   string
	to     []string
	ymAddr string
}

// Login will authenticate the user. You can use any authentication system here
// `state` is a struct with IP addr, etc. So we can use the YM API service in order
// to request environment / entity services about this IP
func (b *backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username != "username" || password != "password" {
		return nil, errors.New("invalid credentials")
	}
	log.Println("successful authentication: ", state.RemoteAddr.String())
	return &session{
		to:     make([]string, 0),
		ymAddr: b.ymAddr,
	}, nil
}

// AnonymousLogin is the handler for anynmous auth
func (b *backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

// Mail will fetch the sender
func (s *session) Mail(from string) error {
	s.from = from
	return nil
}

// Rcpt will fetch the recipients
func (s *session) Rcpt(to string) error {
	s.to = append(s.to, to)
	return nil
}

// Data will return the whole email
func (s *session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		// send the actual to YM
		if err := smtp.SendMail(s.ymAddr, nil, s.from, s.to, bytes.NewReader(b)); err != nil {
			return err
		}
		// reset the session
		s.Reset()
	}
	return nil
}

// Reset will discard the currently processed message
func (s *session) Reset() {
	s.from = ""
	s.to = nil
}

// Logout will free all resources associated with session
func (s *session) Logout() error { return nil }

// StartProxy starts the mail proxy in order
// to add a middleware layer between YM and internet
func StartProxy(proxyConf string) error {
	conf, err := ioutil.ReadFile(proxyConf)
	if err != nil {
		return err
	}
	var c Configuration
	if err := yaml.Unmarshal(conf, &c); err != nil {
		return err
	}
	s := smtp.NewServer(&backend{ymAddr: c.YMAddr})
	s.Addr = fmt.Sprintf("%s:%d", c.Addr, c.Port)
	s.MaxRecipients = c.MaxRecipients
	s.AllowInsecureAuth = true
	log.Println("starting proxy on: ", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
