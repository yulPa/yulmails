package api

import "fmt"

// Configuration is the global configuration of the API server
type Configuration struct {
	// Database configuration
	Database ConfDB `json:"database"`
	// Server configuration
	Server ConfSrv `json:"server"`
}

// ConfDB is the database configuration
type ConfDB struct {
	// Username of the database
	Username string `json:"username"`
	// Password of the database
	Password string `json:"password"`
	// Host of the database
	Host string `json:"host"`
	// Port of the database
	Port int `json:"port"`
	// Name of the database to use
	Name string `json:"name"`
}

type ConfSrv struct {
	// Port to listen on
	Port int `json:"port"`
}

func StartAPI(apiConfig string) error {
	fmt.Println("config: ", apiConfig)
	return nil
}
