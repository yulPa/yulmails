package domains

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateANewAuthorizedPool(t *testing.T) {
	e := []Environment{
		Environment{
			Name: "environment1",
			domains: []string{
				"domain1",
				"domain2",
				"domain3",
			},
		},
		Environment{
			Name: "environment2",
			domains: []string{
				"domain1",
				"domain2",
				"domain3",
			},
		},
	}
	a := NewAuthorizedPool(e)
	assert.Equal(t, "domain1", a.Env[1].domains[0])
}

func TestCreateAAuthorizedPoolFromJson(t *testing.T) {
	_data := []byte(`
    {
      "environments": [
        {
          "name": "zimbra",
          "domains": [
            "zimbra1.com",
            "zimbra2.com",
            "zimbraN.com"
          ]
        },
        {
          "name": "http",
          "domains": [
            "http1.com",
            "http2.com",
            "httpN.com"
          ]
        }
      ]
    }
  `)
	a := CreateANewAuthorizedPoolFromJson(_data)
	assert.Equal(t, a.Env[1].Name, "http")
}
