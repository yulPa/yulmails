package main

import (
	"github.com/yulPa/yulmails/api"
	"github.com/yulPa/yulmails/configuration"
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/mongo"
	"github.com/yulPa/yulmails/sender"
)

var log = logger.GetLogger()

func main() {

	// var workdb = mongo.NewSession("mongodb://workdb:27017")
	var archivingdb = mongo.NewSession("mongodb://archivingdb:27017")
	defer archivingdb.Close()

	err := configuration.NewConfigurationFromConfFile(archivingdb)

	if err != nil {
		log.Errorln(err)
		panic(err)
	}

	/*
		Start Go Subroutines
	*/
	// go entrypoint.Run()

	go sender.Run()
	api.Start()
}
