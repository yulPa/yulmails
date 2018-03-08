package api

import (
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/mongo"

	"net/http"
)

func Start(certFile string, keyFile string) {
	/* This method will start v1 API server */

	var log = logger.GetLogger()
	var archivingdb = mongo.NewSession("mongodb://archivingdb:27017")
	var workdb = mongo.NewSession("mongodb://archivingdb:27017")

	log.Info("Start server for API V1")
	log.Errorln(http.ListenAndServeTLS(":443", certFile, keyFile, GetRouterV1(archivingdb)))
	log.Errorln(http.ListenAndServe(":9252", GetDockerRouterV1(workdb)))
}
