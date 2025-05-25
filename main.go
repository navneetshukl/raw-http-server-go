package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on PORT", port)

	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("error in binding the port:", err)
		os.Exit(1)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("error in accepting the connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error in reading data from connection:", err)
		os.Exit(1)
	}

	data := string(buffer[:n])
	fmt.Println("Received request:\n", data)
	fmt.Println("******************************************************************")
	reqArray := strings.Split(data, "\r\n")

	for _, line := range reqArray {
		fmt.Println(line)
	}
	fmt.Println("******************************************************************")

	lines := strings.Split(data, "\r\n")
	if len(lines) == 0 || !strings.HasPrefix(lines[0], "GET") {
		fmt.Println("Invalid request")
		return
	}

	path := strings.Split(lines[0], " ")[1]
	body := "Requested path: " + path + "\n"

	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
		len(body), body,
	)

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("error in writing response:", err)
	}
}
