package proxy

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration is the configuration for 
// the mail proxy.
type Configuration struct {
	// Port of the proxy
	Port int `yaml:"port"`
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
	fmt.Println(c.Port)
	return nil
}
