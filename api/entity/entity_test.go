package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type entityMock struct{}

func (e *entityMock) ListEntity() ([]*entity, error) {
	return []*entity{
		&entity{
			Id:          1,
			Name:        "entity-1",
			Created:     "2019-01-25 13:34:32",
			Description: "the first entity",
		},
	}, nil
}

func TestGetEntities(t *testing.T) {
	repo := &entityMock{}
	entities, _ := repo.ListEntity()
	assert.Equal(t, len(entities), 1)
}
