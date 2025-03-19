package websockets

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allows cross-origin WebSocket connections
	},
}

var clients = make(map[string]*websocket.Conn)
var mu sync.Mutex

// Handle WebSocket Connections
func HandleConnections(c *gin.Context) {
	username := c.Query("username")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WebSocket Upgrade Error:", err)
		return
	}

	// Store WebSocket Connection
	mu.Lock()
	clients[username] = conn
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, username)
		mu.Unlock()
		conn.Close()
	}()

	// Listen for messages (if needed)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

// Notify a user when they are assigned a task
func NotifyUser(username, message string) {
	mu.Lock()
	conn, ok := clients[username]
	mu.Unlock()

	if ok {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
