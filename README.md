# TCP Echo Server

A simple TCP server implementation in Go that listens for incoming client connections and echoes back received messages.

## Overview

This project demonstrates a basic TCP server that:
- Listens on port 3000 for incoming TCP connections
- Accepts multiple concurrent client connections
- Reads messages from each client
- Echoes back the received messages with an "Echo: " prefix
- Handles client disconnections gracefully

## Project Structure

```
tcp_go/
├── go.mod      # Go module definition
├── main.go     # Main server implementation
└── README.md   # This file
```

## Requirements

- **Go Version**: 1.25.2 or higher
- **Port**: 3000 (must be available)

## Code Breakdown

### Main Components

#### `main()` Function
The entry point of the application that:
1. Creates a TCP listener on port 3000 using `net.Listen()`
2. Logs an error and exits if the server fails to start
3. Defers the listener closure to ensure proper cleanup
4. Enters an infinite loop to accept incoming client connections
5. Spawns a new goroutine for each accepted connection to handle it concurrently

```go
listener, err := net.Listen("tcp", PORT)
```

#### `handleConnection()` Function
Manages individual client connections:
1. Receives a `net.Conn` object representing the client connection
2. Defers connection closure when the function returns
3. Creates a buffered reader to read data from the client
4. Continuously reads messages from the client (one line at a time, terminated by `\n`)
5. Echoes back the message with an "Echo: " prefix
6. Returns gracefully when the client disconnects (EOF error)
7. Logs the client's remote address when disconnected

## Usage

### Starting the Server

```bash
go run main.go
```

You should see:
```
TCP Server started on localhost
```

### Testing the Server

You can test the server using various methods:

#### Using `nc` (netcat):
```bash
nc localhost 3000
```

Then type a message and press Enter:
```
Hello Server
```

Expected response:
```
Echo: Hello Server
```

#### Using `telnet`:
```bash
telnet localhost 3000
```

#### Using `curl`:
```bash
echo "Hello" | nc localhost 3000
```

## Key Features

- **Concurrent Connections**: Uses goroutines to handle multiple clients simultaneously
- **Graceful Shutdown**: Properly closes listener and connections
- **Error Handling**: Catches and logs connection errors without crashing
- **Simple Protocol**: Echo-based protocol for easy testing and understanding

## How It Works

1. Server starts and listens on `localhost:3000`
2. When a client connects, the server accepts the connection
3. A new goroutine is spawned to handle the client independently
4. The handler reads messages line-by-line from the client
5. Each message is echoed back with a prefix
6. When the client closes the connection, the handler gracefully terminates

## Constants

- **PORT**: `:3000` - The port on which the server listens

## Potential Improvements

- Make the port configurable via command-line flags or environment variables
- Add connection timeouts to prevent hanging connections
- Implement a more sophisticated protocol beyond simple echoing
- Add graceful shutdown handling (SIGINT/SIGTERM signals)
- Implement connection pooling or rate limiting
- Add more detailed logging with timestamps

## Error Handling

The server handles two main error scenarios:
1. **Server startup failure**: Logs the error and exits with status code 1
2. **Connection acceptance failure**: Logs the error and continues accepting new connections
3. **Client disconnection**: Logs the client address and gracefully closes the handler

## License

This project is provided as-is for educational purposes.
