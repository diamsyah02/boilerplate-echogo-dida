package auth_test

import (
	"boilerplate-echogo-dida/modules/auth"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockService struct {
	loginFunc    func(user auth.Users) (auth.Users, bool, error)
	registerFunc func(user auth.Users) error
}

func (m *mockService) Login(user auth.Users) (auth.Users, bool, error) {
	return m.loginFunc(user)
}
func (m *mockService) Register(user auth.Users) error {
	return m.registerFunc(user)
}

func TestLogin(t *testing.T) {
	// Persiapkan mock repository
	mockRepo := &mockService{}
	service := auth.NewAuthService(mockRepo)

	// Data uji
	user := auth.Users{Username: "testuser", Password: "testpass"}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	mockUser := auth.Users{Username: "testuser", Password: string(hashedPassword)}

	// Mock repositori
	mockRepo.loginFunc = func(user auth.Users) (auth.Users, bool, error) {
		return mockUser, true, nil
	}

	// Testing Login dengan password yang benar
	token, err := service.Login(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Testing Login dengan password yang salah
	user.Password = "wrongpass"
	token, err = service.Login(user)
	assert.NoError(t, err)
	assert.Empty(t, token)

	// Testing user tidak ditemukan
	mockRepo.loginFunc = func(user auth.Users) (auth.Users, bool, error) {
		return auth.Users{}, false, nil
	}
	token, err = service.Login(user)
	assert.NoError(t, err)
	assert.Empty(t, token)
}

func TestRegister(t *testing.T) {
	// Persiapkan mock repository
	mockRepo := &mockService{}
	service := auth.NewAuthService(mockRepo)

	// Data uji
	user := auth.Users{Username: "newuser", Password: "newpass"}

	// Mock register
	mockRepo.registerFunc = func(user auth.Users) error {
		return nil
	}

	// Testing Register
	err := service.Register(user)
	assert.NoError(t, err)

	// Test case ketika ada error pada repository
	mockRepo.registerFunc = func(user auth.Users) error {
		return errors.New("error")
	}
	err = service.Register(user)
	assert.Error(t, err)
}
