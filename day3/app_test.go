package day3

import (
	"freshers-bootcamp/database"
	mock_database "freshers-bootcamp/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Using Mock Gen
func TestGetUserNameByIDManualMocking(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	mockDB := &database.MockDatabase{}
	mockDB.On("GetUserByID", 1).Return(database.User{Name: "Alice"}, nil)

	name, err := GetUserNameByID(mockDB, 1)

	assert.Nil(t, err)
	assert.Equal(t, "Alice", name)
}

// Using Mock Gen
func TestGetUserNameByID(t *testing.T) {

	//mockgen -destination=mocks/mock_interface.go github.com/self/freshers-bootcamp/database  Database
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_database.NewMockDatabase(ctrl)
	mockDB.EXPECT().GetUserByID(1).Return(database.User{Name: "Alice"}, nil)

	name, err := GetUserNameByID(mockDB, 1)

	assert.Nil(t, err)
	assert.Equal(t, "Alice", name)
}
