package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "192.168.0.5:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// Start a loop to continuously read and write data to the server
	for {
		// Read input from the user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// Send the message to the server
		conn.Write([]byte(message))

		// Read the response from the server
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		// Print the response from the server
		fmt.Println("Received response:", string(buf))
	}
}
