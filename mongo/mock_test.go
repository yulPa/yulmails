package mongo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAMongoInstance(t *testing.T) {
	conn := NewMockSession()
	assert.IsType(t, conn, MockSession{})
}

func TestGetADatabase(t *testing.T) {
	conn := NewMockSession()
	db := conn.DB("a_database")
	assert.IsType(t, db, MockDatabase{})
}

func TestGetACollection(t *testing.T) {
	conn := NewMockSession()
	db := conn.DB("a_database")
	col := db.C("a_collection")
	assert.IsType(t, col, MockCollection{})
}

func TestGetCountOnCollection(t *testing.T) {
	conn := NewMockSession()
	db := conn.DB("a_database")
	col := db.C("a_collection")
	count, _ := col.Count()
	assert.Equal(t, count, 10)
}

func TestCopyASession(t *testing.T) {
	conn := NewMockSession()
	copyConn := conn.Copy()
	assert.IsType(t, copyConn, MockSession{})
}

func TestInsertSomething(t *testing.T)  {

	type Test struct {Yolo string}

	conn := NewMockSession()
	db := conn.DB("a_database")
	col := db.C("a_collection")
	err := col.Insert(&Test{Yolo: "hello"})
	assert.Nil(t, err)
}

func TestFindSomething(t *testing.T)  {

	type Test struct {Yolo string}

	conn := NewMockSession()
	db := conn.DB("a_database")
	col := db.C("a_collection")
	res := col.Find(&Test{Yolo: "hello"})
	assert.IsType(t, res, MockQuery{})
}

func TestFetchSomething(t *testing.T)  {

	type Test struct {Yolo string}
	var test Test

	conn := NewMockSession()
	db := conn.DB("a_database")
	col := db.C("a_collection")
	res := col.Find(&Test{Yolo: "hello"}).All(test)

	assert.Nil(t, res)
	assert.IsType(t, test, Test{})

}
