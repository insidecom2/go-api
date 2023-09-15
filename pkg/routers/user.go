package routers

import (
	"demoecho/pkg/controllers"
	"demoecho/pkg/services"

	"github.com/labstack/echo/v4"
)

var (
	userService     services.UserService       = services.NewUserService()
	userControllers controllers.UserController = controllers.NewUserController(userService)
)

type Router interface {
	InitUserRoute(e *echo.Echo)
}

func InitUserRoute(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users/:id", userControllers.GetUser)
		v1.POST("/users", userControllers.CreateUser)
	}
}
