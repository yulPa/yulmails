package mongo

import "fmt"

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
