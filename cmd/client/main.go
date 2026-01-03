package main

import (
	"fmt"
	"net"
)

const PORT string = ":3000"

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	conn.Write([]byte("Hello, Server"))

	buffer := make([]byte, 1024)

	n, _ := conn.Read(buffer)

	/* if err != nil {
		log.Println("Error reading the message")
	} */

	fmt.Println("Server reply:", string(buffer[:n]))

}
