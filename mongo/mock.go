package mongo

import "fmt"

type MockSession struct{}
type MockDatabase struct {
	Name string
}
type MockCollection struct {
	FullName string
}

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

func (md MockDatabase) C(name string) Collection {
	return MockCollection{
		FullName: fmt.Sprintf("%s.%s", md.Name, name),
	}
}

func (mc MockCollection) Count() (n int, err error) {
	return 10, nil
}
