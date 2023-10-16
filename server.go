package main

import (
	"demoecho/pkg/database"
	"demoecho/pkg/middlewares"
	"demoecho/pkg/routers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	err := godotenv.Load(".env")
	var path = "/api/v1"
	if err != nil {
		log.Panic("Cannot get .env file")
	}
	// connect DB //
	database.ConnectDB()

	// migration //
	if os.Getenv("ENV") == "development" {
		database.AutoMigration()
	}

	// logger
	e.Use(middleware.Logger())

	// routes //
	// auth
	auth := e.Group(fmt.Sprintf("%s/auth", path))
	routers.InitAuthRoute(auth)
	// user
	user := e.Group(fmt.Sprintf("%s/user", path))
	user.Use(middlewares.Authorization)
	routers.InitUserRoute(user)

	e.Logger.Fatal(e.Start(":8001"))
}
