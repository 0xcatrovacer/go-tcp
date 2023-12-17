package client

import (
	"fmt"
	"net"
)

func ReceiveResponse(connection net.Conn) (string, error) {
	response := make([]byte, 1024)
	n, err := connection.Read(response)

	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	return string(response[:n]), nil
}
