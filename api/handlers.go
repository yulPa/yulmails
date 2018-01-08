package api

import (
	"github.com/yulPa/yulmails/domains"
	"github.com/yulPa/yulmails/logger"

	"io/ioutil"
	"net/http"
)

var log = logger.GetLogger()

func CreateAuthorizedPool(w http.ResponseWriter, r *http.Request) {
	/*
	   Create an authorized Pool from HTTP request
	*/

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	pool := domains.CreateANewAuthorizedPoolFromJson(b)
	log.Infoln(pool)
}
