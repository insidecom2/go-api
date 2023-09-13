package main

import (
	"demoecho/pkg/db"
	"demoecho/pkg/models"
	"demoecho/pkg/routers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Cannot get .env file")
	}
	// connect DB //
	db.ConnectDB()

	// migration //
	models.AutoMigration()

	// routes
	routers.InitUserRoute(e)

	e.Logger.Fatal(e.Start(":8001"))
}
