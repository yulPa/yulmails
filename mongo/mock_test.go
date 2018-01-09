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
