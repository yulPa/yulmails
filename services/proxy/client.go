package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"

	"github.com/yulpa/yulmails/api/domain"
)

// Yulmails is the programmatic API
// to use with the proxy
// later we will move elsewhere to be globally
// available
type Yulmails interface {
	// GetDomain return a domain, nil if not found
	GetDomain(string) (*domain.Domain, error)
	// GetWhitelist return IP whitelisted for YM
	GetWhitelist() ([]string, error)
}

type client struct {
	api string
	h   *http.Client
	c   *cache.Cache
}

// yulmails is the http address of yulmails api server
func newClient(yulmails string) *client {
	// add custom auth etc. here
	return &client{
		api: yulmails,
		h:   &http.Client{},
		// 5min of TTL / 10 min before being purged
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *client) GetDomain(name string) (*domain.Domain, error) {
	resp, err := c.h.Get(fmt.Sprintf("%s/domains", c.api))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get domains")
	}
	if resp.StatusCode != 200 {
		// todo returns the error if existing
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to extract body")
	}
	var ds []domain.Domain
	if err := json.Unmarshal(body, &ds); err != nil {
		return nil, errors.Wrapf(err, "unable to read JSON")
	}
	// this shit has to be rework once the
	// filter API will be implemented
	for _, d := range ds {
		if d.Name == name {
			return &d, nil
		}
	}
	return nil, nil
}

func (c *client) GetWhitelist() ([]string, error) {
	if ips, ok := c.c.Get("whitelist"); ok {
		log.Printf("using cached whitelist\n")
		return ips.([]string), nil
	}
	resp, err := c.h.Get(fmt.Sprintf("%s/whitelist", c.api))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get whitelist")
	}
	if resp.StatusCode != 200 {
		// todo returns the error if existing
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to extract body")
	}
	var ips []string
	if err := json.Unmarshal(body, &ips); err != nil {
		return nil, errors.Wrapf(err, "unable to read JSON")
	}
	c.c.Set("whitelist", ips, cache.DefaultExpiration)
	return ips, nil
}
