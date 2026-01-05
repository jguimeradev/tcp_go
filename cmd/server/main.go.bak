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

	log.Println("TCP Server started on localhost")

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)

	}

}

var (
	clientsPool = make(map[net.Conn]string)
	messages    = make(chan string, 10)
	clientsMu   = make(chan struct{}, 1) // mutex for clientsPool
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	// Add client to pool
	clientsMu <- struct{}{}
	clientsPool[conn] = conn.RemoteAddr().String()
	<-clientsMu

	// welcome message
	_, err := conn.Write([]byte("Welcome to the server\n"))

	if err != nil {
		log.Printf("Server write error: %v", err)
	}

	// read messages from this client
	go func() {
		reader := bufio.NewReader(conn)

		for {
			// server reads the message from client
			message, err := reader.ReadString('\n')

			// when client disconnect
			if err != nil {
				log.Println("Client disconnected:", conn.RemoteAddr())
				// Remove client from pool
				clientsMu <- struct{}{}
				delete(clientsPool, conn)
				<-clientsMu
				return
			}

			messages <- fmt.Sprintf("Client %s: %s", conn.RemoteAddr().String(), message)
		}
	}()
}

func init() {
	// broadcast messages to all clients
	go func() {
		for message := range messages {
			clientsMu <- struct{}{}
			for client := range clientsPool {
				_, err := client.Write([]byte(message))
				if err != nil {
					log.Printf("Error writing to client: %v", err)
				}
			}
			<-clientsMu
		}
	}()
}
