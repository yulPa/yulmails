package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateANewEntity(t *testing.T) {
	conservation := OptsConservation{
		sent:   3,
		unsent: 1,
		keep:   true,
	}

	e := NewEntity(
		"An Entity",
		"entity@toto.com",
		conservation,
	)

	assert.Equal(t, e.Conservation.sent, 3)
	assert.True(t, e.Conservation.keep)
}

func TestCreateANewEntityWithoutKeepParameter(t *testing.T) {
	conservation := OptsConservation{
		sent:   3,
		unsent: 1,
	}

	e := NewEntity(
		"An Entity",
		"entity@toto.com",
		conservation,
	)

	assert.False(t, e.Conservation.keep)
}
