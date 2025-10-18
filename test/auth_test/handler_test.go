package auth_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"boilerplate-echogo-dida/modules/auth"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	loginFunc    func(user auth.Users) (string, error)
	registerFunc func(user auth.Users) error
}

func (m *mockHandler) Login(user auth.Users) (string, error) {
	return m.loginFunc(user)
}
func (m *mockHandler) Register(user auth.Users) error {
	return m.registerFunc(user)
}

func TestLogin_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockHandler{
		loginFunc: func(user auth.Users) (string, error) {
			return "TOKEN123", nil
		},
	}

	h := auth.NewAuthHandler(mock)
	err := h.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "TOKEN123")
}

func TestLoginHandler_EmptyFields(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"","password":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := auth.NewAuthHandler(&mockHandler{})
	err := h.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Username and password is required")
}

func TestRegister_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"username":"dida","password":"123"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockHandler{
		registerFunc: func(user auth.Users) error {
			return nil
		},
	}

	h := auth.NewAuthHandler(mock)
	err := h.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Register success")
}

func TestRegister_EmptyFields(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"username":"","password":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := auth.NewAuthHandler(&mockHandler{})
	err := h.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Username and password is required")
}
