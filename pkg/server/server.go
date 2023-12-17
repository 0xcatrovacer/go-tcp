package server

import (
	"fmt"
	"log"
	"net"
)

func HandleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := connection.Read(buffer)

		if err != nil {
			log.Println("Error reading from connection: ", err)
			return
		}

		data := buffer[:n]
		fmt.Printf("Received: %s\n", string(data))

		_, writeErr := connection.Write(data)
		if writeErr != nil {
			log.Println("Error writing to connection: ", writeErr)
		}
	}
}
