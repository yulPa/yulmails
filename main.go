package main

import (
	"fmt"

	"github.com/yulPa/yulmails/api"
	"github.com/yulPa/yulmails/configuration"
	"github.com/yulPa/yulmails/entrypoint"
)

func main() {
	conf := configuration.NewConfigurationFromConfFile()
	fmt.Println(conf)

	/*
		Start Go Subroutines
	*/
	go entrypoint.Run()

	api.Start()
}
