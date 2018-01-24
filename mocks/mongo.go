package mocks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	pb "github.com/yulPa/yulmails/api/mongopb/proto"
	"github.com/yulPa/yulmails/entity"
	"github.com/yulPa/yulmails/environment"
	"github.com/yulPa/yulmails/mongo"
)

type MockSession struct{}
type MockDatabase struct {
	Name string
}
type MockCollection struct {
	FullName string
}
type MockQuery struct{}

func NewMockSession() mongo.Session {
	return MockSession{}
}

func (ms MockSession) Close() {}

func (ms MockSession) DB(name string) mongo.DataLayer {
	mockDatabase := MockDatabase{
		Name: name,
	}
	return mockDatabase
}

func (ms MockSession) Copy() mongo.Session {
	return ms
}

func (md MockDatabase) C(name string) mongo.Collection {
	return MockCollection{
		FullName: fmt.Sprintf("%s.%s", md.Name, name),
	}
}

func (mc MockCollection) Count() (n int, err error) {
	return 10, nil
}

func (mc MockCollection) Find(query interface{}) mongo.Query {
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

func (mc MockCollection) Remove(result interface{}) error {
	return nil
}

func (mc MockCollection) Update(selector interface{}, update interface{}) error {
	return nil
}

func (md MockDatabase) ReadEntities() ([]entity.Entity, error) {
	absPath, _ := filepath.Abs("../mocks/fixtures/entity/entities.json")
	data, _ := ioutil.ReadFile(absPath)
	return entity.NewEntities(data), nil
}

func (md MockDatabase) ReadEntity(name string) (*entity.Entity, error) {
	if name == "an_entity" {
		absPath, _ := filepath.Abs("../mocks/fixtures/entity/entity.json")
		data, _ := ioutil.ReadFile(absPath)
		ent, _ := entity.NewEntity(data)
		return ent, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (md MockDatabase) CreateEntity(ent []byte) error {
	_, err := entity.NewEntity(ent)
	if err != nil {
		return err
	}
	return nil
}

func (md MockDatabase) CreateEnvironment(ent string, env []byte) error {
	_, err := environment.NewEnvironment(env)
	if err != nil {
		return err
	}
	return nil

}

func (md MockDatabase) ReadEnvironment(entName string, envName string) (*environment.Environment, error) {
	if entName == "an_entity" && envName == "an_environment" {
		absPath, _ := filepath.Abs("../mocks/fixtures/environment/environment.json")
		data, _ := ioutil.ReadFile(absPath)
		env, _ := environment.NewEnvironment(data)
		return env, nil
	}
	return nil, errors.New("not found")
}

func (md MockDatabase) DeleteEntity(entName string) error {
	if entName == "an_entity" {
		return nil
	}
	return errors.New("not found")
}

func (md MockDatabase) UpdateEntity(entName string, ent []byte) error {
	if entName == "an_entity" {
		return nil
	}
	return errors.New("not found")
}

func (md MockDatabase) DeleteEnvironment(entName string, envName string) error {
	if entName == "an_entity" && envName == "an_environment" {
		return nil
	}
	return errors.New("not found")
}

func (md MockDatabase) UpdateEnvironment(entName string, envName string, env []byte) error {
	if entName == "an_entity" && envName == "an_environment" {
		return nil
	}
	return errors.New("not found")
}

func (md MockDatabase) ReadEnvironments(entName string) ([]environment.Environment, error) {
	if entName == "an_entity" {
		absPath, _ := filepath.Abs("../mocks/fixtures/environment/environments.json")
		data, _ := ioutil.ReadFile(absPath)
		env, _ := environment.NewEnvironments(data)
		return env, nil
	}
	return nil, errors.New("not found")
}

func (md MockDatabase) SaveMail(entName string, envName string, mail *pb.MailMessage) error {
	return nil
}

func (md MockDatabase) ReadMails(entName string, envName string) ([]pb.MailMessage, error) {
	return nil, errors.New("not found")
}
