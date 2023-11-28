package routers

import (
	handlers "demoecho/pkg/handlers/user"
	repositories "demoecho/pkg/repositories/user"
	services "demoecho/pkg/services/user"

	"github.com/labstack/echo/v4"
)

var (
	userRepo    repositories.UserRepository = repositories.NewUserRepo()
	userService services.UserService        = services.NewUserService(userRepo)
	userHandler handlers.UserHandler        = handlers.NewUserHandler(userService)
)

type UserRouter interface {
	InitUserRoute(u *echo.Group)
}

func InitUserRoute(u *echo.Group) {
	u.GET("", userHandler.GetUser)
	u.GET("/:id", userHandler.GetUserById)
}
