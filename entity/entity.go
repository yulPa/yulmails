package entity

import "encoding/json"

type Entity struct {
	Name         string           `json:"name"`
	Abuse        string           `json:"abuse"`
	Conservation OptsConservation `json:"conservation"`
}

type OptsConservation struct {
	sent   int  `json:"sent"`
	unsent int  `json:"unsent"`
	keep   bool `json:"keep,omitempty"`
}

type Entitys []Entity

func newEntity(name string, abuse string, conservation OptsConservation) *Entity {
	/*
	   Create a new Entity
	   parameter: <string> Entity name
	   parameter: <string> Default abuse address
	   parameter: <OptsConservation> Options relating to conservation time
	   return: <Entity> Return a new entity
	*/
	return &Entity{
		Name:         name,
		Abuse:        abuse,
		Conservation: conservation,
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
