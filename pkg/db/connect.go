package db

import (
	"demoecho/pkg/configs"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

type Connect interface {
	ConnectDB() *gorm.DB
}

func ConnectDB() *gorm.DB {
	log.Printf("Start connect %s", configs.GetDBDriver())
	db, err := gorm.Open(configs.GetDBDriver(), configs.GetConnectionString())

	if err != nil {
		log.Panicf("Cannot connect DB %s", err)
	}
	log.Printf("Connected %s success", configs.GetDBDriver())

	DB = db
	return db
}
