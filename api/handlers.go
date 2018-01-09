package api

import (
	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/environment"
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/mongo"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var log = logger.GetLogger()

func CreateEntity(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
	   Create an entity Pool from HTTP request
	*/
	// TODO: Insert into DB
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	entity := entity.NewEntity(b)
	// TODO: Send entity into MONGO DB
	log.Info(entity)
}

func ReadEntities(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Return a JSON list of present Entities
	*/

	// TODO: Fetch Entities from DB
	e := []byte(`
		{
			"entities": [
				{
		      "name": "an_entity",
		      "abuse": "abuse@domain.tld",
		      "conservation":{
		        "sent": 5,
		        "unsent": 2,
		        "keep": true
		      }
				},
				{
		      "name": "another_entity",
		      "abuse": "abuse1@domain.tld",
		      "conservation":{
		        "sent": 5,
		        "unsent": 2,
		        "keep": true
		      }
				}
			]
    }
		`)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(e)
}

func CreateEnvironment(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		This method will create an environment associated to an entity
	*/

	// TODO: Insert into DB
	vars := mux.Vars(r)
	entityName := vars["entity"]

	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	env := environment.NewEnvironment(b)
	// TODO: Send entity into MONGO DB
	log.Info(entityName, env)
}
