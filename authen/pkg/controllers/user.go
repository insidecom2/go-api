package controllers

import (
	"demoecho/pkg/response"
	"demoecho/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	userService services.UserService
)

type UserController interface {
	GetUser(c echo.Context) (err error)
}

type controllers struct{}


func NewUserController(services services.UserService) UserController {
	userService = services
	return &controllers{}
}

func (con *controllers) GetUser(c echo.Context) (err error) {
	
	 id := c.Param("id")
	 
	result,err := userService.GetUserService(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail("Not found user"))
	}
	resUser := &response.ResponseUser{
		ID: result.ID,
		Name: result.Name,
		Surname: result.Surname,
		Email: result.Email,
	}
	return c.JSON(http.StatusOK, response.ResponseSuccess("OK",resUser))
}


