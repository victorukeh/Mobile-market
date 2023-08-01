package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/controllers"
)

func ErrandRoutes(router fiber.Router) {
	nested := router.Group("/v1/errands")
	errand := &controllers.ErrandController{}

	nested.Post("/add", errand.AddErrand)
}
