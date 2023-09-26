package routers

import (
	"demoecho/pkg/validators"
	"demoecho/pkg/controllers"
	"demoecho/pkg/repositories"
	"demoecho/pkg/services"

	"github.com/labstack/echo/v4"
)

var (
	
	userRepo     	repositories.UserRepository       	= repositories.NewUserRepo()
	userService     services.UserService       			= services.NewUserService(userRepo)
	userControllers controllers.UserController 			= controllers.NewUserController(userService)
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
