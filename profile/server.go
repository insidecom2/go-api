package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"profile/pkg/database"
	"profile/pkg/repositories"
	"profile/pkg/services"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	profileRepo repositories.ProfileRepo = repositories.NewProfileRepo()
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Cannot get .env file")
	}

	// connect DB //
	database.ConnectDB()

	// migration //
	if os.Getenv("ENV") == "development" {
		database.AutoMigration()
	}

	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterProfileServer(s, services.NewProfileServer(profileRepo))

	fmt.Println("GRPC Profile server start on port 50051")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
