package proxy

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/emersion/go-smtp"
	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
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
	YMAPI string `yaml:"yulmails_api"`
}

type backend struct {
	ymAddr string
	api    Yulmails
}
type session struct {
	from   string
	to     []string
	ymAddr string
}

// Login will authenticate the user. You can use any authentication system here
// `state` is a struct with IP addr, etc. So we can use the YM API service in order
// to request environment / entity services about this IP
func (b *backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	raddr := strings.Split(state.RemoteAddr.String(), ":")[0]
	whitelist, err := b.api.GetWhitelist()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get whitelist")
	}
	if !isWhitelisted(raddr, whitelist) {
		log.Printf("auth failed for %s\n", raddr)
		return nil, nil
	}
	log.Printf("%s is now connected\n", raddr)
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

func isWhitelisted(IP string, whitelist []string) bool {
	for _, i := range whitelist {
		if IP == i {
			return true
		}
	}
	return false
}

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
	s := smtp.NewServer(&backend{
		ymAddr: c.YMAddr,
		api:    newClient(c.YMAPI),
	})
	s.Addr = fmt.Sprintf("%s:%d", c.Addr, c.Port)
	s.MaxRecipients = c.MaxRecipients
	s.AllowInsecureAuth = true
	log.Println("starting proxy on: ", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
