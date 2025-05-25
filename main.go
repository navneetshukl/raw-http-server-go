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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error in accepting the connection:", err)
			continue
		}

		go handleMultipleConnection(conn)
	}
}

func handleMultipleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error in reading data from connection:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println("Received request:\n", data)
	fmt.Println("******************************************************************")

	reqArray := strings.Split(data, "\r\n")
	for _, line := range reqArray {
		fmt.Println(line)
	}
	fmt.Println("******************************************************************")

	if len(reqArray) == 0 || !strings.HasPrefix(reqArray[0], "GET") {
		fmt.Println("Invalid request")
		return
	}

	path := strings.Split(reqArray[0], " ")[1]
	body := "Requested path: " + path + "\n"

	if path == "/" || path == "/index" {
		htmlData, err := loadHTML()
		if err != nil {
			fmt.Println("Error reading HTML file:", err)
			conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
			return
		}

		headers := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/html\r\n" +
			fmt.Sprintf("Content-Length: %d\r\n", len(htmlData)) +
			"\r\n"

		_, err = conn.Write([]byte(headers))
		_, err = conn.Write(htmlData)
		if err != nil {
			fmt.Println("error in writing response:", err)
		}
	} else {
		response := fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(body), body,
		)

		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("error in writing response:", err)
		}
	}
}
