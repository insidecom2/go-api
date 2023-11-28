package auth

import (
	"demoecho/pkg/models"
	"demoecho/pkg/requests"
	"demoecho/pkg/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (con *authHandlers) Login(c echo.Context) (err error) {
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

func (con *authHandlers) Register(c echo.Context) (err error) {
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

func (con *authHandlers) RefreshToken(c echo.Context) (err error) {
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
