package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewConfigurationFromFile(t *testing.T) {
	conf := NewConfigurationFromConfFile()
	assert.Equal(t, conf.DbUser, "superman")
}
