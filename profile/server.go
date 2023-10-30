package main

import (
	"fmt"
	"log"
	"net"
	"profile/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	// e := echo.New()
	// server.RegisterRouterV1(e)
	// e.Logger.Fatal(e.Start(":8002"))

	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterProfileServer(s, services.NewProfileServer())

	fmt.Println("GRPC Profile server start on port 50051")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
