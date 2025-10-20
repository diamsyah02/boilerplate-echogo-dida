package auth_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"boilerplate-echogo-dida/internal/auth"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockHandlerService struct {
	loginFunc    func(user auth.Users) (string, error)
	registerFunc func(user auth.Users) error
}

func (m *mockHandlerService) Login(user auth.Users) (string, error) {
	return m.loginFunc(user)
}
func (m *mockHandlerService) Register(user auth.Users) error {
	return m.registerFunc(user)
}

type mockDistributor struct {
	distributeFunc func(to, subject, body string) error
}

func (m *mockDistributor) DistributeEmail(to, subject, body string) error {
	if m.distributeFunc != nil {
		return m.distributeFunc(to, subject, body)
	}
	return nil
}

func setupContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestLogin_Success(t *testing.T) {
	c, rec := setupContext(http.MethodPost, "/login", `{"username":"dida","password":"123"}`)

	mockSvc := &mockHandlerService{
		loginFunc: func(user auth.Users) (string, error) {
			return "TOKEN123", nil
		},
	}
	mockDist := &mockDistributor{}

	h := auth.NewAuthHandler(mockSvc, mockDist)
	err := h.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "TOKEN123")
}

func TestLogin_EmptyFields(t *testing.T) {
	c, rec := setupContext(http.MethodPost, "/login", `{"username":"","password":""}`)

	mockSvc := &mockHandlerService{}
	mockDist := &mockDistributor{}

	h := auth.NewAuthHandler(mockSvc, mockDist)
	err := h.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Username and password is required")
}

func TestRegister_Success(t *testing.T) {
	c, rec := setupContext(http.MethodPost, "/register", `{"username":"dida","password":"123"}`)

	mockSvc := &mockHandlerService{
		registerFunc: func(user auth.Users) error {
			return nil
		},
	}
	mockDist := &mockDistributor{
		distributeFunc: func(to, subject, body string) error {
			// Bisa tambahkan validasi jika perlu
			return nil
		},
	}

	h := auth.NewAuthHandler(mockSvc, mockDist)
	err := h.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Register success")
}

func TestRegister_EmptyFields(t *testing.T) {
	c, rec := setupContext(http.MethodPost, "/register", `{"username":"","password":""}`)

	mockSvc := &mockHandlerService{}
	mockDist := &mockDistributor{}

	h := auth.NewAuthHandler(mockSvc, mockDist)
	err := h.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Username and password is required")
}

func TestRegister_EmailTaskFails(t *testing.T) {
	c, rec := setupContext(http.MethodPost, "/register", `{"username":"dida","password":"123"}`)

	mockSvc := &mockHandlerService{
		registerFunc: func(user auth.Users) error {
			return nil
		},
	}
	mockDist := &mockDistributor{
		distributeFunc: func(to, subject, body string) error {
			return errors.New("failed to queue email task")
		},
	}

	h := auth.NewAuthHandler(mockSvc, mockDist)
	err := h.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Internal server error")
}
