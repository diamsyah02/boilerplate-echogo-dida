package auth

import (
	"net/http"

	"boilerplate-echogo-dida/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type TaskDistributorInterface interface {
	DistributeEmail(to, subject, body string) error
}

type AuthHandler struct {
	service     AuthService
	distributor TaskDistributorInterface
}

func NewAuthHandler(s AuthService, d TaskDistributorInterface) AuthHandler {
	return AuthHandler{
		service:     s,
		distributor: d,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	logger := log.With().Str("function", "Login").Logger()

	var req Users
	if err := c.Bind(&req); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to bind user")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Cannot read request body", nil))
	}
	if req.Username == "" || req.Password == "" {
		logger.Error().
			Msg("Username and password is required")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username and password is required", nil))
	}
	token, err := h.service.Login(req)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to login")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil))
	}
	if token == "" {
		logger.Warn().
			Msg("Failed to login, username or password is wrong")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username or password is wrong", nil))
	}
	data := map[string]interface{}{
		"username": req.Username,
		"token":    token,
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Login success", data))
}

func (h *AuthHandler) Register(c echo.Context) error {
	logger := log.With().Str("function", "Register").Logger()

	var req Users
	if err := c.Bind(&req); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to bind user")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Cannot read request body", nil))
	}
	if req.Username == "" || req.Password == "" {
		logger.Error().
			Msg("Username and password is required")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username and password is required", nil))
	}
	err := h.service.Register(req)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to register")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil))
	}
	err = h.distributor.DistributeEmail(req.Email, "Welcome to our app", "Thank you for registering")
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to send email")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil))
	}
	data := map[string]interface{}{
		"username": req.Username,
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Register success", data))
}
