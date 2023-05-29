package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/finance"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinanceController struct{}

func (uc *FinanceController) CreateWallet(c *fiber.Ctx) error {
	var wallet models.Wallet
	userID, _ := c.Locals("id").(primitive.ObjectID)
	check, _ := c.Locals("id").(string)
	if len(userID) < 24 {
		fmt.Println("UserId: ", userID)
		fmt.Println("UserId: ", check)
		response := &handler.ErrorResponse{Success: true, Error: "Unauthorized"}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	fmt.Println("UserId: ", userID)
	createWallet := models.Wallet{
		UserID:  userID,
		Balance: 0,
	}
	result, err := wallet.Create(createWallet)
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: err.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	response := &finance.CreateWalletResponse{Success: true, Message: "Wallet Created Successfully", Wallet: &result}
	return c.Status(fiber.StatusCreated).JSON(response)
}
