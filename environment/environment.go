package environment

import (
	"encoding/json"

	"github.com/yulPa/yulmails/options"
)

type Environment struct {
	Name     string          `json:"name"`
	IPs      []string        `json:"ips"`
	Abuse    string          `json:"abuse,omitempty"`
	IsOpen   bool            `json:"open"`
	Options  options.Options `json:"options"`
	IdEntity string          `json:"entity,omitempty"`
}

type Environments []Environment

func NewDefaultEnvironment(name string, ips []string, abuse string, isOpen bool) *Environment {
	/*
			   Create a new default environment with default quota values
			   parameter: <[]string> String arrays of IPs address allowed to send email
			   parameter: <string> Abuse address
			   parameter: <bool> If True, we do not need to provider authorized IPs
		     return: <Environment> A new environment
	*/
	return &Environment{
		Name:   name,
		IPs:    ips,
		Abuse:  abuse,
		IsOpen: isOpen,
	}
}

func NewEnvironment(data []byte) *Environment {
	/*
	   Create a new environment directly from a JSON struct
	   parameter: <[]byte> Environment Json array
	   return: <Environment> A new environment
	*/
	var env Environment
	json.Unmarshal(data, &env)
	return &env
}
