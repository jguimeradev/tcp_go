# Changes to main.go.bak - TCP Server Message Broadcasting

## Issues Fixed

### 1. **Unused Connection Pool Channel**
- **Problem**: `connectionsPool` was created as a channel, sent to once, and received once, serving no purpose and wasting a goroutine.
- **Fix**: Removed the `connectionsPool` channel entirely.

### 2. **Local Scope ClientsPool**
- **Problem**: `clientsPool` was created locally within `handleConnection()`, meaning each client connection had its own isolated map with only one client. Messages couldn't be broadcast to multiple clients.
- **Fix**: Moved `clientsPool` to a global variable so all connections share the same client registry.

### 3. **Message Goroutine Exits After First Message**
- **Problem**: The message broadcasting goroutine only read one message (`message := <-messages`) and then exited, preventing any further message distribution.
- **Fix**: Wrapped the message reading in an infinite `for` loop within the `init()` function to continuously process messages.

### 4. **No Continuous Broadcasting**
- **Problem**: Messages were never continuously broadcast to all clients due to the single-read issue.
- **Fix**: Implemented a dedicated `init()` function that launches a goroutine to continuously read from the `messages` channel and send to all connected clients.

### 5. **Race Condition on ClientsPool**
- **Problem**: Multiple goroutines were accessing `clientsPool` without synchronization, causing potential data corruption.
- **Fix**: Implemented a channel-based mutex (`clientsMu`) to safely guard access to `clientsPool`.

### 6. **No Client Cleanup**
- **Problem**: Disconnected clients remained in `clientsPool`, causing write errors on broken connections.
- **Fix**: Added cleanup logic that removes clients from the pool when they disconnect.

### 7. **Missing Error Handling on Write**
- **Problem**: Write errors to clients were silently ignored.
- **Fix**: Added error checking and logging when writing to clients.

## Implementation Details

### Global Variables
```go
var (
	clientsPool = make(map[net.Conn]string)     // All connected clients
	messages    = make(chan string, 10)          // Buffered message queue
	clientsMu   = make(chan struct{}, 1)         // Synchronization for clientsPool
)
```

### Key Functions

**`handleConnection(conn net.Conn)`**
- Adds new client to the global `clientsPool`
- Sends welcome message
- Spawns a goroutine to read messages from this specific client
- Removes client from pool on disconnect

**`init()`**
- Runs automatically on startup
- Spawns a goroutine that continuously reads from the `messages` channel
- Broadcasts each message to all connected clients with error handling

## How It Works

1. Server listens for incoming connections
2. Each new connection is added to `clientsPool`
3. A dedicated goroutine reads messages from each connected client
4. Messages are sent to the shared `messages` channel
5. The broadcast goroutine continuously reads messages and sends them to all clients
6. When a client disconnects, it's removed from `clientsPool`
