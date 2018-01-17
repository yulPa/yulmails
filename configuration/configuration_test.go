package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewConfigurationFromFile(t *testing.T) {
	conf := NewConfigurationFromConfFile()
	assert.Equal(t, "1", conf.V)
	assert.Equal(t, "archiving_db", conf.S.Archiving.Name)
}
