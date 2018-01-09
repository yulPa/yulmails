package environment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateANewEnvironment(t *testing.T) {
	data := []byte(`
    {
      "ips": [
        "192.168.0.1",
        "192.168.0.2",
        "192.168.0.3"
      ],
      "abuse": "abuse@domain.tld",
      "open": false,
      "quota": {
        "tenlastminutes": 150,
        "sixtylastminutes": 200,
        "lastday": 1000,
        "lastweek": 3000,
        "lastmonth": 10000
      }
    }
    `)

	env := NewEnvironment(data)
	assert.Equal(t, 3000, env.Quota.LastWeek)
}
