package routers

import (
	"demoecho/pkg/controllers"
	"demoecho/pkg/repositories"
	"demoecho/pkg/services"
	"demoecho/pkg/validators"

	"github.com/labstack/echo/v4"
)

var (
	authValidator 	 validators.Validator				= validators.NewValidate()
	authRepo     	repositories.AuthRepository       	= repositories.NewAuthRepo()
	authService     services.AuthService       			= services.NewAuthService(authRepo)
	authControllers controllers.AuthController 			= controllers.NewAuthController(authService,authValidator)
)

type AuthRouter interface {
	InitAuthRoute(e *echo.Echo)
}

func InitAuthRoute(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/register", authControllers.Register)
		v1.POST("/login", authControllers.Login)
	}
}
