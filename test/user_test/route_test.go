package user_test

import (
	"boilerplate-echogo-dida/modules/user"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockUserRoute struct {
	getUsersFunc   func() ([]user.Users, error)
	updateUserFunc func(user.Users) error
}

// GetUsers implements user handler.
func (m *mockUserRoute) GetUsers() ([]user.Users, error) {
	return m.getUsersFunc()
}

// UpdateUser implements user handler.
func (m *mockUserRoute) UpdateUser(user user.Users) error {
	return m.updateUserFunc(user)
}

func TestGetUsersRoute(t *testing.T) {
	e := echo.New()
	mockHandler := &mockUserRoute{
		getUsersFunc: func() ([]user.Users, error) {
			return []user.Users{}, nil
		},
	}
	userHandler := user.NewUserHandler(mockHandler)
	api := e.Group("/api")
	api.GET("/users", userHandler.GetUsers)

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Get users success")
}

func TestUpdateUserRoute(t *testing.T) {
	e := echo.New()

	mockHandler := &mockUserRoute{
		updateUserFunc: func(user user.Users) error {
			return nil
		},
	}
	userHandler := user.NewUserHandler(mockHandler)
	api := e.Group("/api")
	api.PUT("/users", userHandler.UpdateUser)

	req := httptest.NewRequest(http.MethodPut, "/api/users", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Update user success")
}
