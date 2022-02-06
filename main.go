package main

import (
	"context"
	"fmt"

	"github.com/nmeisenzahl/grpc-health/private/server"
)

func main() {
	serverAddress := "0.0.0.0:3000"

	server, err := server.New(serverAddress)
	if err != nil {
		fmt.Print("cannot create server: " + err.Error())
	}
	server.Start(context.TODO())
}
