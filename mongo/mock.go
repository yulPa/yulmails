package mongo

type MockSession struct{}
type MockDatabase struct{}
type MockCollection struct{}

func NewMockSession() Session {
	return MockSession{}
}

func (ms MockSession) Close() {}

func (ms MockSession) DB(name string) DataLayer {
	mockDatabase := MockDatabase{}
	return mockDatabase
}

func (md MockDatabase) C(name string) Collection {
	return MockCollection{}
}

func (mc MockCollection) Count() (n int, err error) {
	return 10, nil
}
