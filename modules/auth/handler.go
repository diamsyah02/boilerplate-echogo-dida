package auth

import (
	"net/http"

	"boilerplate-echogo-dida/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func LoginHandler(c echo.Context) error {
	logger := log.With().Str("function", "LoginHandler").Logger()

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
	token, err := LoginService(req)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to login")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil))
	}
	data := map[string]interface{}{
		"username": req.Username,
		"token":    token,
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Login success", data))
}

func RegisterHandler(c echo.Context) error {
	logger := log.With().Str("function", "RegisterHandler").Logger()

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
	err := RegisterService(req)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to register")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil))
	}
	data := map[string]interface{}{
		"username": req.Username,
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Register success", data))
}
