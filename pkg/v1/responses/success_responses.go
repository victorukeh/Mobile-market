package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/finance"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

func WalletSuccessResponse(c *fiber.Ctx, findWallet models.Wallet) error {
	response := &finance.CreateWalletDto{Success: true, Message: "Wallet Found", Wallet: &findWallet}
	return c.Status(fiber.StatusOK).JSON(response)
}
