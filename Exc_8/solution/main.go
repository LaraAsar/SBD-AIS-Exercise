package main

import (
	"exc8/client"
	"exc8/server"
	"time"
)

func main() {
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(1 * time.Second)

	// todo start client
	grpcClient, err := client.NewGrpcClient()
	if err != nil {
		panic(err)
	}

	if err := grpcClient.Run(); err != nil {
		panic(err)
	}
}
