package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/finance"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

func CashTypeCreateResponse(c *fiber.Ctx, result models.CashType) error {
	response := &finance.CreateCashTypeDto{Success: true, Message: "Cash Type Creation Successful", CashType: &result}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func CashGroupCreateResponse(c *fiber.Ctx, result []models.CashGroup) error {
	response := &finance.CreateCashGroupDto{Success: true, Message: "Creation of Cash Groups Successful", CashGroup: &result}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func ErrandCreationResponse(c *fiber.Ctx) error {
	response := &finance.CreateWallet{Success: true, Message: "Errand added successfully"}
	return c.Status(fiber.StatusCreated).JSON(response)
}
