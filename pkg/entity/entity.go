package entity

import (
	"encoding/json"

	"github.com/yulPa/yulmails/pkg/logger"
	"github.com/yulPa/yulmails/pkg/options"
)

var log = logger.GetLogger("entity-ym")

type Entity struct {
	Name    string          `json:"name"`
	Abuse   string          `json:"abuse"`
	Options options.Options `json:"options"`
}

func newEntity(name string, abuse string, opts options.Options) *Entity {
	/*
	   Create a new Entity
	   parameter: <string> Entity name
	   parameter: <string> Default abuse address
	   parameter: <options.Options> Options relating
	   return: <Entity> Return a new entity
	*/
	return &Entity{
		Name:    name,
		Abuse:   abuse,
		Options: opts,
	}
}

func NewEntity(data []byte) (*Entity, error) {
	/*
	   Create a new Entity from a json struct
	   parameter: <[]byte> Json struct
	   return: <Entity> Return a new entity
	*/
	var entity Entity
	err := json.Unmarshal(data, &entity)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &entity, nil
}

func NewEntities(data []byte) []Entity {
	/*
		Create a new array of entity fron a JSON array struct
		parameter: <[]byte> JSON entity array
		return: <[]entity.Entity> An array of entities
	*/
	entities := make([]Entity, 0)
	json.Unmarshal(data, &entities)
	return entities
}
