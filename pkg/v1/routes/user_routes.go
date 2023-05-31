package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/controllers"
	"github.com/victorukeh/mobile-market/pkg/v1/middleware"
)

func UserRoutes(router fiber.Router) {
	// Create a new instance of Fiber for your nested routes
	nested := router.Group("/v1/users", middleware.Authentication)
	user := &controllers.UserController{}
	// Define your nested routes using the nested.Get(), nested.Post(), nested.Put(), and nested.Delete() methods
	nested.Get("/", user.GetUsers)
	nested.Get("/:id", user.GetUser)
}
