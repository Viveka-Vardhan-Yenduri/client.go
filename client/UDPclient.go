package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	message := []byte("Hello from UDP client!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return
	}

	response := string(buffer[:n])
	fmt.Println("Received response:", response)
}
