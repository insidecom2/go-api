package routers

import (
	"demoecho/pkg/controllers"
	"demoecho/pkg/repositories"
	"demoecho/pkg/services"

	"github.com/labstack/echo/v4"
)

var (
	userRepo        repositories.UserRepository = repositories.NewUserRepo()
	userService     services.UserService        = services.NewUserService(userRepo)
	userControllers controllers.UserController  = controllers.NewUserController(userService)
)

type UserRouter interface {
	InitUserRoute(u *echo.Group)
}

func InitUserRoute(u *echo.Group) {
	u.GET("/:id", userControllers.GetUser)
}
