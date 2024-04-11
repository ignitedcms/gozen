package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[*websocket.Conn]bool)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Add new client to the map
	clients[conn] = true

	// Handle WebSocket messages here
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients, conn) // Remove the client from the map
			return
		}

		fmt.Printf("Received message: %s\n", data)

		// Broadcast the message to all connected clients
		for client := range clients {
			if err := client.WriteMessage(messageType, data); err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client) // Remove the client from the map
			}
		}
	}
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.StripPrefix("/", fs))

	// Handle WebSocket requests
	http.HandleFunc("/ws", handleWebSocket)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

