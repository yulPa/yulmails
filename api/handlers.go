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
	db := sess.DB("configuration")

	e, err := db.ReadEntities()

	raw, _ := json.Marshal(e)

	if err != nil {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func ReadEntity(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Read one entity following entity name in url
	*/
	vars := mux.Vars(r)
	entityName := vars["entity"]

	// Open collection associated
	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	ent, err := db.ReadEntity(entityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		raw, _ := json.Marshal(ent)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	}
}

func ReadEnvironment(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Read one environment following entity name and environment name in url
	*/
	vars := mux.Vars(r)
	entityName := vars["entity"]
	environmentName := vars["environment"]

	// Open collection associated
	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	env, err := db.ReadEnvironment(entityName, environmentName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		raw, _ := json.Marshal(env)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	}
}

func DeleteEntity(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Update an entity already in database
	*/
	vars := mux.Vars(r)
	entityName := vars["entity"]

	// Open collection associated
	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	err := db.DeleteEntity(entityName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateEntity(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Update an existing entity
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

	err := db.UpdateEntity(entityName, b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteEnvironment(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Delete an environment
	*/
	vars := mux.Vars(r)
	entName := vars["entity"]
	envName := vars["environment"]

	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	err := db.DeleteEnvironment(entName, envName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateEnvironment(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Update an existing environment
	*/
	vars := mux.Vars(r)
	entName := vars["entity"]
	envName := vars["environment"]

	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	// Read raw content
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	err := db.UpdateEnvironment(entName, envName, b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func ReadEnvironments(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Read all environments associated to an entity
	*/
	vars := mux.Vars(r)
	entName := vars["entity"]

	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	envs, err := db.ReadEnvironments(entName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		raw, _ := json.Marshal(envs)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	}
}

func ReadMails(session mongo.Session, w http.ResponseWriter, r *http.Request) {
	/*
		Read all stored mails associated to an environment
	*/
	vars := mux.Vars(r)
	entName := vars["entity"]
	envName := vars["environment"]

	sess := session.Copy()
	db := sess.DB("configuration")
	defer sess.Close()

	mails, err := db.ReadMails(entName, envName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		raw, _ := json.Marshal(mails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	}
}
