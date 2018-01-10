package entity

import (
	"encoding/json"

	"github.com/yulPa/yulmails/options"
)

type Entity struct {
	Name    string `json:"name"`
	Abuse   string `json:"abuse"`
	Options options.Options `json:"options"`
}

type Entities struct {
	List []Entity `json:"entities"`
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

func NewEntity(data []byte) *Entity {
	/*
	   Create a new Entity from a json struct
	   parameter: <[]byte> Json struct
	   return: <Entity> Return a new entity
	*/
	var entity Entity
	json.Unmarshal(data, &entity)
	return &entity
}

func NewEntities(data []byte) *Entities {
	/*
	 Create new entities from a json struct
	 parameter: <[]byte> Json struct
	 return: <Entites> Return a new Entites sruct
	*/
	var entities Entities
	json.Unmarshal(data, &entities)
	return &entities
}
