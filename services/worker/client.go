package worker

import (
	"encoding/json"
	"net/http"
	"bytes"

	"github.com/yulpa/yulmails/plugins/sdk"
)

type plugin struct {
	client     *http.Client
	pluginAddr string
}

// SendEmail will send an email to the plugin
func (p *plugin) SendEmail(email string) (*sdk.Result, error) {
	payload, err := json.Marshal(map[string]string{
		"email": email,
	})
	if err != nil {
		return nil, err
	}
	// TODO: protocol must be customizable
	resp, err := p.client.Post(
		"http://"+p.pluginAddr+"/check",
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}
	var result sdk.Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPlugin returns a plugin created from pluginAddr
// and with worker configuration (TLS, etc.)
func NewPlugin(pluginAddr string) *plugin {
	return &plugin{
		client:     &http.Client{},
		pluginAddr: pluginAddr,
	}
}
