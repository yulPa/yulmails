package entity

import (
	"encoding/json"

	"github.com/yulPa/yulmails/options"
)

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
