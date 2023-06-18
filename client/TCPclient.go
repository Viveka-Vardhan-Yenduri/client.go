package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to TCP server:", err)
		return
	}
	defer conn.Close()

	message := "Hello from TCP client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return
	}

	response := string(buffer)
	fmt.Println("Received response:", response)
}
