package controllers

import (
	"demoecho/pkg/models"
	"demoecho/pkg/requests"
	"demoecho/pkg/response"
	"demoecho/pkg/services"
	"demoecho/pkg/validators"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(c echo.Context) (err error)
	Register(c echo.Context) (err error)
	RefreshToken(c echo.Context) (err error)
}

var (
	authService      services.AuthService
	authValidatorReq validators.Validator
)

type authController struct{}

func NewAuthController(services services.AuthService, validator validators.Validator) AuthController {
	authService = services
	authValidatorReq = validator
	return &authController{}
}

func (con *authController) Login(c echo.Context) (err error) {
	login := new(requests.LoginBody)
	v := authValidatorReq.Validate(c, login)

	if v != nil {
		return v
	}

	token, refreshToken, err := authService.LoginAuth(*login)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ResponseFail(err.Error()))
	}

	res := &response.LoginAuthRes{
		TOKEN:         token,
		REFRESH_TOKEN: refreshToken,
	}

	return c.JSON(http.StatusOK, response.ResponseSuccess("OK", res))
}

func (con *authController) Register(c echo.Context) (err error) {
	user := new(models.User)

	v := authValidatorReq.Validate(c, user)

	if v != nil {
		return v
	}

	result, err := authService.RegisterAuth(*user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail(err.Error()))
	}

	resUser := &response.ResponseUser{
		ID:      result.ID,
		Name:    result.Name,
		Surname: result.Surname,
		Email:   result.Email,
	}

	return c.JSON(http.StatusOK, response.ResponseSuccess("OK", resUser))
}

func (con *authController) RefreshToken(c echo.Context) (err error) {
	email := fmt.Sprint(c.Get("email"))
	if email == "" {
		return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid Token"))
	}

	token, refreshToken, err := authService.RefreshAuth(email)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ResponseFail(err.Error()))
	}

	res := &response.LoginAuthRes{
		TOKEN:         token,
		REFRESH_TOKEN: refreshToken,
	}

	return c.JSON(http.StatusOK, response.ResponseSuccess("OK", res))
}
