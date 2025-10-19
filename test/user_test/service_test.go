package user_test

import (
	"boilerplate-echogo-dida/internal/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockService struct {
	getUsersFunc   func() ([]user.Users, error)
	updateUserFunc func(user.Users) error
}

func (m *mockService) GetUsers() ([]user.Users, error) {
	return m.getUsersFunc()
}

func (m *mockService) UpdateUser(user user.Users) error {
	return m.updateUserFunc(user)
}

func TestGetUsersService(t *testing.T) {
	mockRepo := &mockService{}
	service := user.NewUserService(mockRepo)

	mockRepo.getUsersFunc = func() ([]user.Users, error) {
		return []user.Users{}, nil
	}

	users, err := service.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, []user.Users{}, users)
}

func TestUpdateUserService(t *testing.T) {
	mockRepo := &mockService{}
	service := user.NewUserService(mockRepo)

	mockRepo.updateUserFunc = func(user user.Users) error {
		return nil
	}

	err := service.UpdateUser(user.Users{})
	assert.NoError(t, err)
}
