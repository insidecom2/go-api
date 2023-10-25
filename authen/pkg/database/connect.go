package database

import (
	"demoecho/pkg/configs"
	"demoecho/pkg/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

type Connect interface {
	ConnectDB()
	AutoMigration()
}

func ConnectDB() {
	log.Printf("Start connect %s", configs.GetDBDriver())
	var err error
	DB, err = gorm.Open(configs.GetDBDriver(), configs.GetConnectionString())

	if err != nil {
		log.Panicf("Cannot connect DB %s", err)
	}
	log.Printf("Connected %s success", configs.GetDBDriver())

}

func AutoMigration() {
	log.Print("Starting Migration.")
	DB.AutoMigrate(&models.User{})
}
