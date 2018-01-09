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

	// Open collection associated
	sess := session.Copy()
	defer sess.Close()
	db := session.DB("configuration")
	col := db.C("entity")

	// Read raw content
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// Create and push the struct
	entity := entity.NewEntity(b)
	err := col.Insert(entity)
	if err != nil {
		log.Error(err)
	}
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
	vars := mux.Vars(r)
	entityName := vars["entity"]

	// Open collection associated
	sess := session.Copy()
	defer sess.Close()
	db := session.DB("configuration")
	col := db.C("environment")

	// Read raw content
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// Create and push the struct
	env := environment.NewEnvironment(b)
	env.IdEntity = entityName
	err := col.Insert(env)
	if err != nil {
		log.Error(err)
	}

	w.WriteHeader(http.StatusCreated)
}
