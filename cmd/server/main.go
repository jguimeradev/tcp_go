package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const PORT string = ":3000"

func main() {
	ln, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Println("Error starting server:", err)
		os.Exit(1)
	}

	defer ln.Close()

	fmt.Println("TCP Server started on localhost")

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}

		fmt.Printf("Received from %v: %s", conn.RemoteAddr(), message)

		conn.Write([]byte("Echo: " + message))

	}
}
