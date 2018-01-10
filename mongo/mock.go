package mongo

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/environment"
)

type MockSession struct{}
type MockDatabase struct {
	Name string
}
type MockCollection struct {
	FullName string
}
type MockQuery struct{}

func NewMockSession() Session {
	return MockSession{}
}

func (ms MockSession) Close() {}

func (ms MockSession) DB(name string) DataLayer {
	mockDatabase := MockDatabase{
		Name: name,
	}
	return mockDatabase
}

func (ms MockSession) Copy() Session {
	return ms
}

func (md MockDatabase) C(name string) Collection {
	return MockCollection{
		FullName: fmt.Sprintf("%s.%s", md.Name, name),
	}
}

func (mc MockCollection) Count() (n int, err error) {
	return 10, nil
}

func (mc MockCollection) Find(query interface{}) Query {
	return MockQuery{}
}

func (mc MockCollection) Insert(docs ...interface{}) error {
	return nil
}

func (mq MockQuery) All(result interface{}) error {
	return nil
}

func (mq MockQuery) One(result interface{}) error {
	return nil
}

func (md MockDatabase) ReadEntities() ([]entity.Entity, error) {
	absPath, _ := filepath.Abs("../mongo/fixtures/entity/entities.json")
	data, _ := ioutil.ReadFile(absPath)
	return entity.NewEntities(data), nil
}

func (md MockDatabase) ReadEntity(name string) (*entity.Entity, error) {
	absPath, _ := filepath.Abs("../mongo/fixtures/entity/entity.json")
	data, _ := ioutil.ReadFile(absPath)
	return entity.NewEntity(data), nil
}

func (md MockDatabase) CreateEntity(ent []byte) error {
	return nil
}

func (md MockDatabase) CreateEnvironment(ent string, env []byte) error {
	return nil
}

func (md MockDatabase) ReadEnvironment(entName string, envName string) (*environment.Environment, error) {
	return &environment.Environment{}, nil
}
