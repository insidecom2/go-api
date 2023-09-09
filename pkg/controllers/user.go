package controllers

import (
	"demoecho/pkg/interfaces"
	"demoecho/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	userService services.UserService
)

type UserController interface {
	GetUser(c echo.Context) (err error)
	CreateUser(c echo.Context) (err error)
}

type controllers struct{}

func NewUserController(services services.UserService) UserController {
	userService = services
	return &controllers{}

}

func (*controllers) GetUser(c echo.Context) (err error) {
	u := &interfaces.User{
		Name:    c.QueryParam("name"),
		Surname: c.QueryParam("surname"),
	}
	result := userService.GetUserService(u)

	return c.JSON(http.StatusOK, result)

}

func (*controllers) CreateUser(c echo.Context) (err error) {
	u := &interfaces.User{
		Name:    c.FormValue("name"),
		Surname: c.FormValue("surname"),
	}
	result := userService.GetUserService(u)

	return c.JSON(http.StatusOK, result)

}
