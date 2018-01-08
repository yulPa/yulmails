package environment

import "encoding/json"

type Environment struct {
	IPs    []string  `json:"ips"`
	Abuse  string    `json:"abuse,omitempty"`
	IsOpen bool      `json:"open"`
	Quota  OptsQuota `json:"quota,omitempty"`
}

type OptsQuota struct {
	TenLastMinutes   int `json:"tenlastminutes"`
	SixtyLastMinutes int `json:"sixtylastminutes"`
	LastDay          int `json:"lastday"`
	LastWeek         int `json:"lastweek"`
	LastMonth        int `json:"lastmonth"`
}

type Environments []Environment

func NewDefaultEnvironment(ips []string, abuse string, isOpen bool) *Environment {
	/*
			   Create a new default environment with default quota values
			   parameter: <[]string> String arrays of IPs address allowed to send email
			   parameter: <string> Abuse address
			   parameter: <bool> If True, we do not need to provider authorized IPs
		     return: <Environment> A new environment
	*/
	return &Environment{
		IPs:    ips,
		Abuse:  abuse,
		IsOpen: isOpen,
		Quota: OptsQuota{
			TenLastMinutes:   10,
			SixtyLastMinutes: 10,
			LastDay:          10,
			LastWeek:         10,
			LastMonth:        10,
		},
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
