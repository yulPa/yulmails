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
	Find(interface{}) Query
	Insert(...interface{}) error
}

type Query interface {
	All(interface{}) error
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
