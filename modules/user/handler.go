package user

import (
	"boilerplate-echogo-dida/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GetUsersHandler(c echo.Context) error {
	users, err := GetUsersService()
	if err != nil {
		log.Error().
			Str("function", "GetUsersHandler").
			Err(err).
			Msg("Failed to get users")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Get users failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Get users success", users))
}

func UpdateUserHandler(c echo.Context) error {
	user := Users{}
	if err := c.Bind(&user); err != nil {
		log.Error().
			Str("function", "UpdateUserHandler").
			Err(err).
			Msg("Failed to bind user")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Update user failed", nil))
	}
	if err := UpdateUserService(user); err != nil {
		log.Error().
			Str("function", "UpdateUserHandler").
			Err(err).
			Msg("Failed to update user")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Update user failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Update user success", nil))
}
