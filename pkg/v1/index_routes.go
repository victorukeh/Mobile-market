package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/middleware"
	"github.com/victorukeh/mobile-market/pkg/v1/routes"
)

func SetupRoutes(app *fiber.App) {
	// Define your top-level routes using the app.Get(), app.Post(), app.Put(), and app.Delete() methods
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About page")
	})

	nested := app.Group("/api")
	// Set up nested routes defined in users_routes.go
	routes.AuthRoutes(nested)

	app.Use(middleware.Authentication)

	routes.UserRoutes(nested)
	routes.FinanceRoutes(nested)
}
