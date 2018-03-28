package main

import (
	"github.com/yulPa/yulmails/pkg/mongo"
	"github.com/yulPa/yulmails/pkg/logger"
)

var (
	log      = logger.GetLogger("yulmails-compute")
)

func main() {

	var workdb = mongo.NewSession("mongodb://workdb:27017")
	log.Infoln(workdb)

}
