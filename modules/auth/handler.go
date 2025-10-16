package auth

import (
	"net/http"

	"boilerplate-echogo-dida/utils"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	var req Users
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Cannot read request body", nil, nil))
	}
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username and password is required", nil, nil))
	}
	token, err := LoginService(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil, err))
	}
	data := map[string]interface{}{
		"username": req.Username,
		"token":    token,
	}
	res := utils.WebResponse(http.StatusOK, "Login success", data, nil)
	return c.JSON(http.StatusOK, res)
}

func RegisterHandler(c echo.Context) error {
	var req Users
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Cannot read request body", nil, nil))
	}
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Username and password is required", nil, nil))
	}
	result, err := RegisterService(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Internal server error", nil, err))
	}
	data := map[string]interface{}{
		"username": result.Username,
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Register success", data, nil))
}
