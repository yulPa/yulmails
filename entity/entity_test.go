package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/yulPa/yulmails/options"
)

func TestCreateANewEntity(t *testing.T) {
	opts := options.Options{
		Conservation: options.OptsConservation{
			Sent:   3,
			Unsent: 1,
			Keep:   true,
		},
	}

	e := newEntity(
		"An Entity",
		"entity@toto.com",
		opts,
	)

	assert.Equal(t, e.Options.Conservation.Sent, 3)
	assert.True(t, e.Options.Conservation.Keep)
}

func TestCreateANewEntityWithoutKeepParameter(t *testing.T) {
	opts := options.Options{
		Conservation: options.OptsConservation{
			Sent:   3,
			Unsent: 1,
		},
	}

	e := newEntity(
		"An Entity",
		"entity@toto.com",
		opts,
	)

	assert.False(t, e.Options.Conservation.Keep)
}

func TestCreateANewEntityFromJson(t *testing.T) {
	_data := []byte(`
    {
      "name": "An entity",
      "abuse": "abuse@domain.tld",
      "options": {
				"conservation":{
	        "sent": 5,
	        "unsent": 2,
	        "keep": true
	      }
			}
    }
    `)

	e := NewEntity(_data)

	assert.Equal(t, "abuse@domain.tld", e.Abuse)
	assert.Equal(t, 2, e.Options.Conservation.Unsent)

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
