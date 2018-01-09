package mongo

import (
	"github.com/yulPa/yulmails/logger"
	"gopkg.in/mgo.v2"
)

var log = logger.GetLogger()

// Create interface Session -> DataLayer -> Collection

type Session interface {
	DB(name string) DataLayer
	Close()
}

type DataLayer interface {
	C(name string) Collection
}

type Collection interface {
	Count() (int, error)
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

func (ms MongoSession) DB(name string) DataLayer {
	/*
	   Return an interface DataLayer which wraps DB object
	   return: <MongoDatabase> Mongo database
	*/
	return &MongoDatabase{
		Database: ms.Session.DB(name),
	}
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

func NewSession() Session {
	/*
	   In this method, we create a new Mongo session in order to dial with database
	   return: <Session> A mongo db session
	*/
	mgoSession, err := mgo.Dial("mongo:27017")
	if err != nil {
		log.Error(err)
		return nil
	}
	return MongoSession{mgoSession}
}
