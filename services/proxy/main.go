package proxy

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
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
	// Enable tls : true/false
	EnableTLS bool `yaml:"enable_tls"`
	// TLS cert path X509
	TLSCertPath string `yaml:"tls_cert_path"`
	// TLS cert path X509
	TLSKeyPath string `yaml:"tls_key_path"`	
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

	// Reverse DNS lookup
	var client_ReverseDns = "unknown"
	
	if reverse, _ := net.LookupAddr(state.RemoteAddr.String()); len(reverse) != 0 {
		client_ReverseDns = reverse[0]
	}

	// client_ip, port, err
	client_ip, _, _ := net.SplitHostPort(state.RemoteAddr.String())

	// TODO: IP_REPUTATION (bd, stats, first time hold , blacklist, etc)
	if client_ip == "172.18.0.1" {
		return nil, errors.New("Sorry, can't talking to you. Please visit https://yulmails.io/dnsbl")
	}
	// Check Username / Password
	if username != "username@domain.com" || password != "password" {

		log.Println("client=",client_ReverseDns,"[", client_ip,"],sasl_method=PLAIN,sasl_username=",username,",failed: authentication failure")

		return nil, errors.New("invalid credentials")
	}



	log.Println("client=",client_ReverseDns,"[", client_ip,"],sasl_method=PLAIN,sasl_username=",username)
	
	// Offer the SMTP SESSION
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
	
	proxyCert, err := tls.LoadX509KeyPair(c.TLSCertPath, c.TLSKeyPath)

    if err != nil {
        log.Fatalf("Error: loadkeys: %s", err)
    }

    if c.EnableTLS {
		s.TLSConfig = &tls.Config{
						Certificates: []tls.Certificate{proxyCert},
					}
		log.Println("starting SMTP proxy with TLS_ENABLE on:", s.Addr)
		log.Println("loading certificate: ", c.TLSCertPath)
	
	} else {
	
		log.Println("starting SMTP proxy (warning : no TLS) on:", s.Addr)
	}
	
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
