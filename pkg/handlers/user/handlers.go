package user

import (
	services "demoecho/pkg/services/user"

	"github.com/labstack/echo/v4"
)

var (
	userService services.UserService
)

type UserHandler interface {
	GetUserById(c echo.Context) (err error)
	GetUser(c echo.Context) (err error)
}

type handlers struct{}

func NewUserHandler(services services.UserService) UserHandler {
	userService = services
	return &handlers{}
}
