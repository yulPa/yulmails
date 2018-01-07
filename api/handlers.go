package api

import (
	"github.com/yulPa/check_mails/logger"
	"github.com/yulPa/check_mails/domains"

	"net/http"
	"io/ioutil"
)

var log = logger.GetLogger()

func CreateAuthorizedPool(w http.ResponseWriter, r *http.Request) {
  /*
    Create an authorized Pool from HTTP request
  */

	b, _ := ioutil.ReadAll(r.Body)
  defer r.Body.Close()

  pool := domains.CreateANewAuthorizedPoolFromJson(b)
}
