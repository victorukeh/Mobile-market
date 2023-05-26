package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// "github.com/gofiber/websocket/v2"
	routes "github.com/victorukeh/mobile-market/pkg/v1"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	// port := os.Getenv("PORT")
	// Create a new instance of Fiber
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// Define your top-level routes using the app.Get(), app.Post(), app.Put(), and app.Delete() methods
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Use Websockets
	// app.Get("/ws", websocket.New(func(c *websocket.Conn) {
	// 	// Handle WebSocket connection here
	// 	for {
	// 		// Read incoming message from WebSocket
	// 		mt, message, err := c.ReadMessage()
	// 		if err != nil {
	// 			log.Println(err)
	// 			break
	// 		}
	// 		// Log incoming message
	// 		log.Printf("Received message: %s\n", message)
	// 		// Write message back to WebSocket
	// 		err = c.WriteMessage(mt, message)
	// 		if err != nil {
	// 			log.Println(err)
	// 			break
	// 		}
	// 	}
	// }))

	// Call the SetupRoutes() function from the routes package to set up all the routes for your application
	routes.SetupRoutes(app)
	// log.Printf("Listening on :%s...", port)
	// Start the server on port 3000
	// app.Listen(fmt.Sprintf(":%s", port))
	// err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("server closed", reason)
	})

	http.Handle("/socket.io/", server)
	go server.Serve()
	defer server.Close()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	app.Listen(":2000")
}
