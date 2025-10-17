package user

import (
	"boilerplate-echogo-dida/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return UserHandler{service}
}

func (h *UserHandler) GetUsersHandler(c echo.Context) error {
	logger := log.With().Str("function", "GetUsersHandler").Logger()

	users, err := h.service.GetUsersService()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to get users")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Get users failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Get users success", users))
}

func (h *UserHandler) UpdateUserHandler(c echo.Context) error {
	logger := log.With().Str("function", "UpdateUserHandler").Logger()

	user := Users{}
	if err := c.Bind(&user); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to bind user")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Update user failed", nil))
	}
	if err := h.service.UpdateUserService(user); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to update user")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Update user failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Update user success", nil))
}
