package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/controllers"
)

func AuthRoutes(router fiber.Router) {
	// Create a new instance of Fiber for your nested routes
	nested := router.Group("/v1/auth")
	auth := &controllers.AuthController{}
	// Define your nested routes using the nested.Get(), nested.Post(), nested.Put(), and nested.Delete() methods
	nested.Post("/register", auth.Register)
	nested.Post("/login", auth.Login)
}
