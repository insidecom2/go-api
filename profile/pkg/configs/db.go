package configs

import (
	"fmt"
	"os"
)

func GetDBDriver() string {
	return os.Getenv("DBDRIVER")
}

func GetConnectionString() string {

	db := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s TimeZone=%s password=%s",
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBUSER"),
		os.Getenv("DBNAME"),
		os.Getenv("DBSSL"),
		os.Getenv("TIMEZONE"),
		os.Getenv("DBPASSWORD"),
	)

	return db
}
