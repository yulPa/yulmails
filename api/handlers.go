package api

import (
	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/logger"

	"io/ioutil"
	"net/http"
)

var log = logger.GetLogger()

func CreateEntity(w http.ResponseWriter, r *http.Request) {
	/*
	   Create an entity Pool from HTTP request
	*/

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	entity := entity.NewEntity(b)
	// TODO: Send entity into MONGO DB
	log.Infoln(entity)
}

func GetEntity(w http.ResponseWriter, r *http.Request) {
	/*
		Return a JSON list of present entitys
	*/

	// TODO: Fetch entitys from DB
	e := []byte(`
		{
		  "name": "An entity",
		  "abuse": "abuse@domain.tld",
		  "conservation":{
		    "sent": 5,
		    "unsent": 2,
		    "keep": true
		  }
		}
		`)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(e)
}
