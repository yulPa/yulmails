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

func TestCreateEntitiesFromJson(t *testing.T) {
	_data := []byte(`
		[
			{
				"name": "an_entity",
				"abuse": "abuse@domain.tld",
				"options": {
					"conservation":{
						"sent": 5,
						"unsent": 2,
						"keep": true
					}
				}
			},
			{
				"name": "another_entity",
				"abuse": "another_abuse@domain.tld",
				"options": {
					"conservation":{
						"sent": 5,
						"unsent": 3,
						"keep": true
					}
				}
			}
		]
		`)

	e := NewEntities(_data)

	assert.Equal(t, "abuse@domain.tld", e[0].Abuse)
	assert.Equal(t, 3, e[1].Options.Conservation.Unsent)

}
