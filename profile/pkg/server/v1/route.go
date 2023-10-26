package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	Context = "profile"
)

func RegisterRouterV1(c *echo.Echo) {
	var path = "/api/v1"

	h := c.Group(fmt.Sprintf("%s/health", path))
	h.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Fine !!!!")
	})

	p := c.Group(fmt.Sprintf("%s/%s", path, Context))
	p.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "Hello "+id)
	})
}
