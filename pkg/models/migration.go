package models

import (
	"demoecho/pkg/db"
	"log"
)

func AutoMigration() {
	log.Print("Starting Migration.")
	db.DB.AutoMigrate(&User{})
}
