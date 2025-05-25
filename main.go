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
	fmt.Println("Listening on PORT ", port)

	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", port))
	if err != nil {
		fmt.Println("error in binding the port ", err)
		os.Exit(1)
	}

	conn, err := list.Accept()
	if err != nil {
		fmt.Println("error in accepting the connection ", err)
		os.Exit(1)
	}

	buff := make([]byte, 1024)

	_, err = conn.Read(buff)
	if err != nil {
		fmt.Println("error in reading data from connection ", err)
		os.Exit(1)
	}

	fmt.Println("data is ", string(buff))

	fmt.Println("******************************************************************")

	reqArray := strings.Split(string(buff), "\r\n")

	for _, v := range reqArray {
		fmt.Println(v)
	}
	fmt.Println("******************************************************************")

}
