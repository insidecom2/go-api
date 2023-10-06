package controllers

import (
	"demoecho/pkg/models"
	"demoecho/pkg/response"
	"demoecho/pkg/services"
	"demoecho/pkg/validators"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	userService services.UserService
	validatorReq validators.Validator
)

type UserController interface {
	GetUser(c echo.Context) (err error)
	CreateUser(c echo.Context) (err error)
}

type controllers struct{}


func NewUserController(services services.UserService) UserController {
	validatorReq = validators.NewValidate()
	userService = services
	return &controllers{}
}

func (con *controllers) GetUser(c echo.Context) (err error) {
	
	 id := c.Param("id")
	 
	result,err := userService.GetUserService(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFail("Not found user"))
	}
	resUser := &response.ResponseUser{
		ID: result.ID,
		Name: result.Name,
		Surname: result.Surname,
		Email: result.Email,
	}
	return c.JSON(http.StatusOK, response.ResponseSuccess("OK",resUser))
}

func (con *controllers) CreateUser(c echo.Context) error {
	
	user := new(models.User)

	v := validatorReq.Validate(c,user)

	if v != nil {
		return v
	}
	
	result,err := userService.CreateUserService(*user)
 
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

