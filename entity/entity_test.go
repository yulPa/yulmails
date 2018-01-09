package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateANewEntity(t *testing.T) {
	conservation := OptsConservation{
		Sent:   3,
		Unsent: 1,
		Keep:   true,
	}

	e := newEntity(
		"An Entity",
		"entity@toto.com",
		conservation,
	)

	assert.Equal(t, e.Conservation.Sent, 3)
	assert.True(t, e.Conservation.Keep)
}

func TestCreateANewEntityWithoutKeepParameter(t *testing.T) {
	conservation := OptsConservation{
		Sent:   3,
		Unsent: 1,
	}

	e := newEntity(
		"An Entity",
		"entity@toto.com",
		conservation,
	)

	assert.False(t, e.Conservation.Keep)
}

func TestCreateANewEntityFromJson(t *testing.T) {
	_data := []byte(`
    {
      "name": "An entity",
      "abuse": "abuse@domain.tld",
      "conservation":{
        "sent": 5,
        "unsent": 2,
        "keep": true
      }
    }
    `)

	e := NewEntity(_data)

	assert.Equal(t, "abuse@domain.tld", e.Abuse)
	assert.Equal(t, 2, e.Conservation.Unsent)

}

func TestCreateNewEntitiesFromJson(t *testing.T) {
	_data := []byte(`
    {
			"entities": [
				{
		      "name": "An entity",
		      "abuse": "abuse@domain.tld",
		      "conservation":{
		        "sent": 5,
		        "unsent": 2,
		        "keep": true
		      }
				},
				{
		      "name": "An entity",
		      "abuse": "abuse@domain.tld",
		      "conservation":{
		        "sent": 5,
		        "unsent": 2,
		        "keep": true
		      }
				}
			]
    }
    `)

	e := NewEntities(_data)

	assert.Len(t, e.List, 2)

}
