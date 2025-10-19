package auth_test

import (
	"boilerplate-echogo-dida/internal/auth"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockAuthRoute struct {
	loginFunc    func(user auth.Users) (string, error)
	registerFunc func(user auth.Users) error
}

// Login implements auth service.
func (m *mockAuthRoute) Login(user auth.Users) (string, error) {
	return m.loginFunc(user)
}

// Register implements auth service.
func (m *mockAuthRoute) Register(user auth.Users) error {
	return m.registerFunc(user)
}

func TestLoginRoute(t *testing.T) {
	e := echo.New()
	mockRoute := &mockAuthRoute{
		loginFunc: func(user auth.Users) (string, error) {
			return "TOKEN_123", nil
		},
	}
	mockHandler := auth.NewAuthHandler(mockRoute)
	api := e.Group("/api")
	api.POST("/login", mockHandler.Login)

	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "TOKEN_123")
}

func TestRegisterRoute(t *testing.T) {
	e := echo.New()

	mockRoute := &mockAuthRoute{
		registerFunc: func(user auth.Users) error {
			return nil
		},
	}
	mockHandler := auth.NewAuthHandler(mockRoute)

	api := e.Group("/api")
	api.POST("/register", mockHandler.Register)

	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Register success")
}
