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

func setupContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestGetUsersHandler(t *testing.T) {
	c, rec := setupContext(http.MethodGet, "/api/users", "")

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
	c, rec := setupContext(http.MethodPut, "/api/users", `{"username":"dida","password":"123"}`)

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
