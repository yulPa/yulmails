package domains

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateANewAllowed(t *testing.T) {
  e := Environment{
    Name: "environment1",
    domains: []string{
      "domain1",
      "domain2",
      "domain3",
    },
  }
  a := NewAllowed(e)
  assert.Equal(t, "domain2", a.Env.domains[1])
}
