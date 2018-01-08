package api

import (
	"github.com/yulPa/yulmails/logger"

	"net/http"
)

func Start() {
	/* This method will start v1 API server */

	var log = logger.GetLogger()

	log.Infoln("Start server for API V1")
	log.Errorln(http.ListenAndServe(":80", GetRouterV1()))

}
