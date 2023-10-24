package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var path = "/api/v1"

	h := e.Group(fmt.Sprintf("%s/health", path))
	h.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Fine !!!!")
	})
	e.Logger.Fatal(e.Start(":8002"))
}
