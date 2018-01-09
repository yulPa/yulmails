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

type Entities struct {
	List []Entity `json:"entities"`
}

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
