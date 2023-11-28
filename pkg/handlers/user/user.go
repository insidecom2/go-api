package user

import (
	"demoecho/pkg/response"
	"demoecho/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (con *handlers) GetUserById(c echo.Context) (err error) {

	id := c.Param("id")

	result, err := userService.GetUserService(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail("Not found user"))
	}
	resUser := &response.ResponseUser{
		ID:      result.ID,
		Name:    result.Name,
		Surname: result.Surname,
		Email:   result.Email,
	}
	return c.JSON(http.StatusOK, response.ResponseSuccess("OK", resUser))
}

func (con *handlers) GetUser(c echo.Context) (err error) {

	id := utils.GetContext(c, "id")
	result, err := userService.GetUserService(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail("Not found user"))
	}
	resUser := &response.ResponseUser{
		ID:      result.ID,
		Name:    result.Name,
		Surname: result.Surname,
		Email:   result.Email,
	}
	return c.JSON(http.StatusOK, response.ResponseSuccess("OK", resUser))
}
