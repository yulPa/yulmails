package main

import (
	"flag"

	"github.com/yulPa/yulmails/api"
	"github.com/yulPa/yulmails/pkg/configuration"
	"github.com/yulPa/yulmails/pkg/logger"
	"github.com/yulPa/yulmails/pkg/mongo"
	"github.com/yulPa/yulmails/pkg/sender"
	"github.com/yulPa/yulmails/pkg/entrypoint"
)

var (
	log      = logger.GetLogger()
	certFile = flag.String("tls-crt-file", "domain.tld.crt", "A certificate file")
	keyFile  = flag.String("tls-key-file", "domain.tld.key", "A key file")
)

func main() {

	flag.Parse()


	var workdb = mongo.NewSession("mongodb://workdb:27017")
	var archivingdb = mongo.NewSession("mongodb://archivingdb:27017")

	log.Info(workdb)
	log.Info(archivingdb)

	defer archivingdb.Close()
	defer workdb.Close()

	err := configuration.NewConfigurationFromConfFile(archivingdb)

	if err != nil {
		log.Errorln(err)
		panic(err)
	}


	/*
		Start Go Subroutines
	*/
	go entrypoint.Run()

	go sender.Run()
	api.Start(*certFile, *keyFile)
}
