package api

import (
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/mongo"

	"net/http"
)

func Start() {
	/* This method will start v1 API server */

	var log = logger.GetLogger()
	var session = mongo.NewSession("mongodb://database:27017")

	log.Info("Start server for API V1")
	log.Errorln(http.ListenAndServe(":80", GetRouterV1(session)))

}
