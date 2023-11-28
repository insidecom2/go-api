package auth

import (
	services "demoecho/pkg/services/auth"
	"demoecho/pkg/validators"

	"github.com/labstack/echo/v4"
)

type AuthHandlers interface {
	Login(c echo.Context) (err error)
	Register(c echo.Context) (err error)
	RefreshToken(c echo.Context) (err error)
}

var (
	authService      services.AuthService
	authValidatorReq validators.Validator
)

type authHandlers struct{}

func NewAuthHandlers(services services.AuthService, validator validators.Validator) AuthHandlers {
	authService = services
	authValidatorReq = validator
	return &authHandlers{}
}
