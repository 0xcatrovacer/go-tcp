package main

import (
	"fmt"
	"log"
	"net"

	"go-tcp/pkg/server"
)

func main() {
	address := "127.0.0.1:8080"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}
	defer listener.Close()

	fmt.Printf("Listening on %s\n", address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %v", err)
			continue
		}

		go server.HandleConnection(connection)
	}
}
