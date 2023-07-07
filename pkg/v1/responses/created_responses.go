package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/finance"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

func WalletCreateResponse(c *fiber.Ctx, result models.CashType) error {
	response := &finance.CreateCashTypeDto{Success: true, Message: "Cash Type Creation Successful", CashType: &result}
	return c.Status(fiber.StatusOK).JSON(response)
}
