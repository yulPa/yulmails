package domains

import "encoding/json"

type Environment struct {
	Name    string   `json:"name"`
	domains []string `json:"domains"`
}

type AuthorizedPool struct {
	Env []Environment `json:"environments"`
}

func NewAuthorizedPool(env []Environment) *AuthorizedPool {
	/*
	   This function create an AuthorizedPool with a given set of environments
	   parameter: <environment> A given set of environment
	   return: <AuthorizedPool> A kind of list of authorized domains
	*/
	return &AuthorizedPool{
		Env: env,
	}
}

func CreateANewAuthorizedPoolFromJson(data []byte) *AuthorizedPool {
	var AuthorizedPool AuthorizedPool
	json.Unmarshal(data, &AuthorizedPool)
	return &AuthorizedPool
}
