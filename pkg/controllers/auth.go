package controllers

import (
	"demoecho/pkg/models"
	"demoecho/pkg/response"
	"demoecho/pkg/services"
	"demoecho/pkg/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Login(c echo.Context) (err error)
	Register(c echo.Context) (err error)
}

var (
	authService services.AuthService
	authValidatorReq validators.Validator
)

type authController struct{}


func NewAuthController(services services.AuthService, validator validators.Validator) AuthController {
	authService = services
	authValidatorReq = validator
	return &authController{}
}

func (con *authController) Login(c echo.Context) (err error) {
	login := new(services.LoginBody)
	v := authValidatorReq.Validate(c,login)

	if v != nil {
		return v
	}

	token,err := authService.LoginAuth(*login)
 
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ResponseFail(err.Error()))
	}

	res := &response.LoginAuthRes{
		TOKEN: token,
	}

	return c.JSON(http.StatusOK, response.ResponseSuccess("OK",res))
}

func (con *authController) Register(c echo.Context) (err error) {
	user := new(models.User)

	v := authValidatorReq.Validate(c,user)

	if v != nil {
		return v
	}
	
	result,err := authService.RegisterAuth(*user)
 
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail(err.Error()))
	}

	resUser := &response.ResponseUser{
		ID: result.ID,
		Name: result.Name,
		Surname: result.Surname,
		Email: result.Email,
	}

	return c.JSON(http.StatusOK, response.ResponseSuccess("OK",resUser))
}