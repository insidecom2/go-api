package validators

import (
	"demoecho/pkg/response"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Response struct {
	res response.Validator
}

type Validator interface {
	Validate(c echo.Context, i interface{}) error 
}


func (r *Response) Validate(c echo.Context, i interface{}) error {
	v := validator.New()
	if err := c.Bind(&i); err != nil {
		c.JSON(http.StatusBadRequest, r.res.ResponseValidator(err.Error()))
		return err
	}
	if err := v.Struct(i); err != nil {
		c.JSON(http.StatusBadRequest, r.res.ResponseValidator(err.Error()))
		return err
	}
	return nil;
}

