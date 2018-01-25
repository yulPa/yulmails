package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/yulPa/yulmails/mocks"
)

func TestCreateNewConfigurationFromFile(t *testing.T) {
	err := NewConfigurationFromConfFile(mocks.NewMockSession())
	assert.Nil(t, err)
}
