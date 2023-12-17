package main

import (
	"bufio"
	"fmt"
	"go-tcp/pkg/client"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr := "127.0.0.1:8080"

	connection, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server: ", err)
		return
	}
	defer connection.Close()

	fmt.Println("Connected to server", serverAddr)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if _, err := connection.Write([]byte(message)); err != nil {
			fmt.Println("Error sending message to the server:", err)
			return
		}

		response, err := client.ReceiveResponse(connection)
		if err != nil {
			fmt.Println("Error receiving response from the server:", err)
			return
		}

		fmt.Println("Server response:", response)
	}
}
