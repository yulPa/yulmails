package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"errors"
	"fmt"

	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/environment"
	"github.com/yulPa/yulmails/logger"
	"github.com/yulPa/yulmails/options"
)

var log = logger.GetLogger()

// Create interface Session -> DataLayer -> Collection

type Session interface {
	DB(name string) DataLayer
	Close()
	Copy() Session
}

type DataLayer interface {
	C(name string) Collection
	ReadEntities() ([]entity.Entity, error)
	ReadEntity(name string) (*entity.Entity, error)
	CreateEnvironment(string, []byte) error
	CreateEntity([]byte) error
	ReadEnvironment(string, string) (*environment.Environment, error)
	DeleteEntity(string) error
	UpdateEntity(string, []byte) error
}

type Collection interface {
	Count() (int, error)
	Find(interface{}) Query
	Insert(...interface{}) error
	Remove(interface{}) error
	Update(interface{}, interface{}) error
}

type Query interface {
	All(interface{}) error
	One(interface{}) error
}

type MongoSession struct {
	*mgo.Session
}

type MongoDatabase struct {
	*mgo.Database
}

type MongoCollection struct {
	*mgo.Collection
}

type MongoQuery struct {
	*mgo.Query
}

func (ms MongoSession) DB(name string) DataLayer {
	/*
	   Return an interface DataLayer which wraps DB object
	   return: <MongoDatabase> Mongo database
	*/
	return &MongoDatabase{
		Database: ms.Session.DB(name),
	}
}

func (ms MongoSession) Copy() Session {
	/*
		Return a copy of MongoSession, in order to allow concurrent job
		return: <Session> A new database session
	*/
	return MongoSession{ms.Session.Copy()}
}

func (md MongoDatabase) C(name string) Collection {
	/*
	   Return an interface Collection which wraps Collection object
	   return: <MongoCollection> A Mongo collection
	*/
	return &MongoCollection{
		Collection: md.Database.C(name),
	}
}

func (mc MongoCollection) Find(query interface{}) Query {
	/*
	 Return real Find value from mongo
	 return: <MongoCollection> A Mongo query
	*/
	return MongoQuery{
		Query: mc.Collection.Find(query),
	}
}

func NewSession(url string) Session {
	/*
		   In this method, we create a new Mongo session in order to dial with database
			 parameter: <string> URL to reach database, in this format: mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
		   return: <Session> A mongo db session
	*/
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		log.Error(err)
		return nil
	}
	return MongoSession{mgoSession}
}

func (md MongoDatabase) ReadEntity(name string) (*entity.Entity, error) {
	/*
		Read a specific entity into database
		parameter: <string> ID/Name of the entity
		return: <entity.Entity> Fetched entity
		return: <error> Nil if no error
	*/
	var res entity.Entity

	colEntity := md.C("entity")
	err := colEntity.Find(bson.M{"name": name}).One(&res)
	if err != nil {
		log.Error(err)
		return &entity.Entity{}, err
	}
	return &res, nil
}

func (md MongoDatabase) CreateEnvironment(ent string, env []byte) error {
	/*
		This method will insert a new environment into DB after checked if options are correct
		parameter: <[]byte> environment JSON
		return: <error> Return `nil` if not error occured
	*/
	associatedEntity, err := md.ReadEntity(ent)
	if err != nil {
		log.Error(err)
		return err
	}

	nEnvironment, err := environment.NewEnvironment(env)
	if err != nil {
		return err
	}
	nEnvironment.EntityId = ent

	if (nEnvironment.Options.Quota == options.OptsQuota{} && associatedEntity.Options.Quota == options.OptsQuota{}) {
		return errors.New(
			fmt.Sprintf(
				"Environment: Quota is not setted for %s entity. Please update %s entity or add quota to %s environment",
				ent,
				nEnvironment.Name,
			),
		)
	} else if (nEnvironment.Options.Quota == options.OptsQuota{} && associatedEntity.Options.Quota != options.OptsQuota{}) {
		nEnvironment.Options.Quota = associatedEntity.Options.Quota
	}

	colEnvironment := md.C("environment")
	err = colEnvironment.Insert(nEnvironment)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (md MongoDatabase) CreateEntity(ent []byte) error {
	/*
		Create and push a new entity in database
		parameter: <[]byte> entity JSON
		return: <error> nil if no error
	*/

	nEntity, err := entity.NewEntity(ent)
	if err != nil {
		log.Error(err)
		return err
	}
	colEntity := md.C("entity")

	err = colEntity.Insert(nEntity)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (md MongoDatabase) ReadEntities() ([]entity.Entity, error) {
	/*
		Return all entities listed in DB
		return: <[]entity.Entity> An entites array
		return: <error> Return nil if no errors
	*/
	var res []entity.Entity

	colEntity := md.C("entity")
	err := colEntity.Find(nil).All(&res)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return res, nil
}

func (md MongoDatabase) ReadEnvironment(entName string, envName string) (*environment.Environment, error) {
	/*
		Read one environment in DB
		parameter: <string> Entity name
		parameter: <string> Environment name
		return: <environment> Wanted environment
		return: <error> Nil if no error
	*/
	var res environment.Environment

	colEnvironment := md.C("environment")
	err := colEnvironment.Find(bson.M{"name": envName, "entityid": entName}).One(&res)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &res, nil
}

func (md MongoDatabase) DeleteEntity(entName string) error {
	/*
		Delete one entity from DB
		parameter: <string> Entity name
		return: <error> Nil if no error
	*/
	colEntity := md.C("entity")

	err := colEntity.Remove(bson.M{"name": entName})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (md MongoDatabase) UpdateEntity(entName string, ent []byte) error {
	/*
		Update an existing entity in database
		parameter: <string> Entity name
		return: <error> Nil if no error
	*/
	colEntity := md.C("entity")

	e, err := entity.NewEntity(ent)
	if err != nil {
		log.Error(err)
		return err
	}

	err = colEntity.Update(bson.M{"name": entName}, &e)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
