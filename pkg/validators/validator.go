package validators

import (
	"demoecho/pkg/response"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type validators struct {}

type Validator interface {
	Validate(c echo.Context, i interface{}) error 
	ResponseValidator(err string) []string
}

func NewValidate() Validator {
	return &validators{}
}

func (r *validators) Validate(c echo.Context, i interface{}) error {
	v := validator.New()
	if err := c.Bind(&i); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseReqFail( r.ResponseValidator(err.Error())))
		return err
	}
	if err := v.Struct(i); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseReqFail( r.ResponseValidator(err.Error())))
		return err
	}
	return nil;
}

func (r *validators) ResponseValidator(err string) []string {

	var error []string
	for _,s := range strings.Split(err,"\n"){
		err := strings.Split(s,":")
		re := regexp.MustCompile(`'[^']+'`)
		newStrs := re.FindAllString(err[2], -1)
		error = append(error, newStrs[0]+" is "+newStrs[1])
	}
	return error
}

