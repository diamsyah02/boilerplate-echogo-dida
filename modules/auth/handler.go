package auth

import (
	"net/http"

	"boilerplate-echogo-dida/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func LoginHandler(c echo.Context) error {
	var req Users
	if err := c.Bind(&req); err != nil {
		log.Error().
			Str("function", "LoginHandler").
			Err(err).
			Msg("Failed to bind user")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Cannot read request body", nil))
	}
	if req.Username == "" || req.Password == "" {
		log.Error().
			Str("function", "LoginHandler").
			Msg("Username and password is required")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username and password is required", nil))
	}
	token, err := LoginService(req)
	if err != nil {
		log.Error().
			Str("function", "LoginHandler").
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
	var req Users
	if err := c.Bind(&req); err != nil {
		log.Error().
			Str("function", "RegisterHandler").
			Err(err).
			Msg("Failed to bind user")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Cannot read request body", nil))
	}
	if req.Username == "" || req.Password == "" {
		log.Error().
			Str("function", "RegisterHandler").
			Msg("Username and password is required")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username and password is required", nil))
	}
	err := RegisterService(req)
	if err != nil {
		log.Error().
			Str("function", "RegisterHandler").
			Err(err).
			Msg("Failed to register")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil))
	}
	data := map[string]interface{}{
		"username": req.Username,
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Register success", data))
}
