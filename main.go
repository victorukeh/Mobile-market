package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	routes "github.com/victorukeh/mobile-market/pkg/v1"
)

func main() {
	// Create a new instance of Fiber
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
	// Call the SetupRoutes() function from the routes package to set up all the routes for your application
	routes.SetupRoutes(app)
	// Start the server on port 3000
	app.Listen(":2000")
}
