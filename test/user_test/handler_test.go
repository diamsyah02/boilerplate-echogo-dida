package user_test

import (
	"boilerplate-echogo-dida/internal/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	getUsersFunc   func() ([]user.Users, error)
	updateUserFunc func(user.Users) error
}

// GetUsers implements user handler.
func (m *mockHandler) GetUsers() ([]user.Users, error) {
	return m.getUsersFunc()
}

// UpdateUser implements user handler.
func (m *mockHandler) UpdateUser(user user.Users) error {
	return m.updateUserFunc(user)
}

func TestGetUsersHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockHandler{
		getUsersFunc: func() ([]user.Users, error) {
			return []user.Users{}, nil
		},
	}
	userHandler := user.NewUserHandler(mock)
	err := userHandler.GetUsers(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Get users success")
}

func TestUpdateUserHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/users", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockHandler{
		updateUserFunc: func(user user.Users) error {
			return nil
		},
	}
	userHandler := user.NewUserHandler(mock)
	err := userHandler.UpdateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Update user success")
}
