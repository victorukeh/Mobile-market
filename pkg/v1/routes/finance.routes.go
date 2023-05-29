package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/controllers"
)

func FinanceRoutes(router fiber.Router){
	nested := router.Group("/v1/finance")
	finance := &controllers.FinanceController{}
	// Define your nested routes using the nested.Get(), nested.Post(), nested.Put(), and nested.Delete() methods
	nested.Post("/create", finance.CreateWallet)
}