package main

import (
	"demoecho/pkg/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routers.InitUserRoute(e)

	e.Logger.Fatal(e.Start(":8001"))
}
