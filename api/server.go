package api

import (
	"github.com/yulPa/check_mails/logger"

	"net/http"
)

func Start() {
	/* This method will start v1 API server */

	var log = logger.GetLogger()

	var routes = Routes{
		Route{
			Name:        "Authorized Pool",
			Method:      http.MethodPost,
			Pattern:     "/api/v1/authorizedpool",
			HandlerFunc: CreateAuthorizedPool,
		},
	}

	var router = NewRouter(routes)

	log.Infoln("Start server for API V1")
	log.Errorln(http.ListenAndServe(":80", router))

}
