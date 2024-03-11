package database

import "github.com/stretchr/testify/mock"

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetUserByID(id int) (User, error) {
	args := m.Called(id)

	if args.Get(1) != nil {
		return User{}, args.Get(1).(error)
	}

	return args.Get(0).(User), nil
}
