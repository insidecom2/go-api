package main

import (
	"profile/pkg/server/v1"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	server.RegisterRouterV1(e)
	e.Logger.Fatal(e.Start(":8002"))
}
