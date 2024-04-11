/*
|---------------------------------------------------------------
| Main entry point
|---------------------------------------------------------------
|
| First load .env variables, register session, csrf middleware
| Load mysql db and finally register routes
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"gozen/db"
	"gozen/routes"
	"gozen/system/rendering"
	"gozen/system/templates"
	"log"
	"net/http"
	"os"
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

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	port := os.Getenv("APP_PORT")

	db.InitDB()

	err = templates.LoadTemplates()
	if err != nil {
		log.Println("Error loading templates:", err)
		return
	}

	// Create a new router instance
	r := chi.NewRouter()

	//r.Use(middleware.StripSlashes)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.CleanPath)

	// Middleware stack
	r.Use(middleware.Logger)

	// Use Gorilla sessions middleware
	r.Use(sessionMiddleware)

	//Use Gorilla CSRF middleware
	r.Use(csrfMiddleware)

	// Handle WebSocket requests
	r.HandleFunc("/ws", handleWebSocket)

	//create an alias to the resources
	//and serve css and js
	r.Handle("/resources/*",
		http.StripPrefix("/resources",
			http.FileServer(http.Dir("./resources/src/assets"))))

	// Custom 404 handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		rendering.RenderTemplate(w, r, "404", nil)
	})

	// Load all routes separately
	routes.LoadRoutes(r)

	foo := `
   ____  ____  ____  ___  ____
  / __  / __ \/_  / / _ \/ __ \
 / /_/ / /_/ / / /_/  __/ / / /
 \__, /\____/ /___/\___/_/ /_/
/____/
`
	fmt.Print(foo)
	fmt.Println("Starting on http://localhost:" + port)
	// Start the HTTP server
	http.ListenAndServe(":"+port, r)

}

// Session middleware
func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session from the request
		// Call the next handler

	   key := os.Getenv("APP_KEY")
      Store := sessions.NewCookieStore([]byte(key))

		Store.Options.HttpOnly = true

		session, _ := Store.Get(r, "session-name")

		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CSRF middleware loads on EVERY route
func csrfMiddleware(next http.Handler) http.Handler {

	//get this from env
	key := os.Getenv("APP_KEY")
   csrfKey := []byte(key)

	csrfMiddleware := csrf.Protect(
		csrfKey,
		csrf.Secure(true), // Set to true for production
	)

	return csrfMiddleware(next)
}
