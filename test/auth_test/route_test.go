package auth_test

import (
	"boilerplate-echogo-dida/modules/auth"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockAuthService struct {
	loginFunc    func(user auth.Users) (string, error)
	registerFunc func(user auth.Users) error
}

func (m *mockAuthService) Login(user auth.Users) (string, error) {
	return m.loginFunc(user)
}

func (m *mockAuthService) Register(user auth.Users) error {
	return m.registerFunc(user)
}

func TestLoginRoute(t *testing.T) {
	e := echo.New()

	mockService := &mockAuthService{
		loginFunc: func(user auth.Users) (string, error) {
			return "TOKEN_123", nil
		},
		registerFunc: func(user auth.Users) error {
			return nil
		},
	}
	mockHandler := auth.NewAuthHandler(mockService)

	api := e.Group("/api")
	api.POST("/login", mockHandler.Login)

	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "TOKEN_123")
}
