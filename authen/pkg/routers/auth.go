package routers

import (
	"demoecho/pkg/controllers"
	"demoecho/pkg/middlewares"
	"demoecho/pkg/repositories"
	"demoecho/pkg/services"
	"demoecho/pkg/validators"

	"github.com/labstack/echo/v4"
)

var (
	authValidator   validators.Validator        = validators.NewValidate()
	authRepo        repositories.AuthRepository = repositories.NewAuthRepo()
	authService     services.AuthService        = services.NewAuthService(authRepo)
	authControllers controllers.AuthController  = controllers.NewAuthController(authService, authValidator)
)

type AuthRouter interface {
	InitAuthRoute(g *echo.Group)
}

func InitAuthRoute(g *echo.Group) {
	g.POST("/register", authControllers.Register)
	g.POST("/login", authControllers.Login)

	g.Use(middlewares.RefreshToken)
	g.GET("/refresh-token", authControllers.RefreshToken)
}
