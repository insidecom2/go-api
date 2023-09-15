package controllers

import (
	"demoecho/pkg/interfaces"
	"demoecho/pkg/response"
	"demoecho/pkg/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var (
	userService services.UserService
)

type UserController interface {
	GetUser(c echo.Context) (err error)
	CreateUser(c echo.Context) (err error)
	Validate(c echo.Context, i interface{}) error 
}

type controllers struct{
}


func NewUserController(services services.UserService) UserController {
	userService = services
	return &controllers{}
}

func (con *controllers) GetUser(c echo.Context) (err error) {
	
	 id := c.Param("id")

	if err != nil {
		return err
	}

	result := userService.GetUserService(id)

	return c.JSON(http.StatusOK, result)
}

func (con *controllers) CreateUser(c echo.Context) error {
	
	user := new(interfaces.User)

	v := con.Validate(c,user)

	if v != nil {
		return v
	}
	
	result := userService.CreateUserService(user)

	resUser := &interfaces.ResponseUser{
		Name: result.Name,
		Surname: result.Surname,
		Email: result.Email,
	}

	return c.JSON(http.StatusOK, resUser)

}

func (con *controllers) Validate(c echo.Context, i interface{}) error {

	v := validator.New()
	if err := c.Bind(&i); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseValidator(err.Error()))
		return err
	}

	if err := v.Struct(i); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseValidator(err.Error()))
		return err
	}

	return nil;
}
