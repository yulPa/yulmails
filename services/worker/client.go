package worker

import (
	"io/ioutil"
	"net/http"
)

type plugin struct {
	client     *http.Client
	pluginAddr string
}

// SendEmail will send an email to the plugin
func (p *plugin) SendEmail(email string) ([]byte, error) {
	// TODO: protocol must be customizable
	resp, err := p.client.Get("http://" + p.pluginAddr + "/check")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// NewPlugin returns a plugin created from pluginAddr
// and with worker configuration (TLS, etc.)
func NewPlugin(pluginAddr string) *plugin {
	return &plugin{
		client:     &http.Client{},
		pluginAddr: pluginAddr,
	}
}
