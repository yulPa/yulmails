package api

import (
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/mongo"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var log = logger.GetLogger()

func CreateEntity(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
	   Create an entity Pool from HTTP request
	*/

	sess := session.Copy()
	defer sess.Close()
	db := session.DB("configuration")

	// Read raw content
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err := db.CreateEntity(b)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)

}

func ReadEntities(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Return a JSON list of present Entities
	*/
	sess := session.Copy()
	defer sess.Close()
	db := sess.DB("entity")

	e, err := db.ReadEntities()

	raw, _ := json.Marshal(e)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	}
}

func CreateEnvironment(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		This method will create an environment associated to an entity
	*/
	vars := mux.Vars(r)
	entityName := vars["entity"]

	// Open collection associated
	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	// Read raw content
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err := db.CreateEnvironment(entityName, b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
