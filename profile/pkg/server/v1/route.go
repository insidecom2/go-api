package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRouterV1(c *echo.Echo) {
	var path = "/api/v1"

	h := c.Group(fmt.Sprintf("%s/health", path))
	h.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Fine !!!!")
	})
}
