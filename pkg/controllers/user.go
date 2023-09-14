package controllers

import (
	"demoecho/pkg/interfaces"
	"demoecho/pkg/services"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
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
	user := &interfaces.User{
		Name:    c.QueryParam("name"),
		Surname: c.QueryParam("surname"),
	}
	result := userService.GetUserService(user)

	return c.JSON(http.StatusOK, result)
}

func (*controllers) CreateUser(c echo.Context) (err error) {
	
	var user interfaces.User
	v := validator.New()
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := v.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, strings.Split(err.Error(),"\n"))
	}

	result := userService.GetUserService(&user)

	resUser := &interfaces.ResponseUser{
		Name: result.Name,
		Surname: result.Surname,
		Email: result.Email,
	}

	return c.JSON(http.StatusOK, resUser)

}
