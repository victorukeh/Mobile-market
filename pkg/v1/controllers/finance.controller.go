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

func (u *FinanceController) GetWallet(c *fiber.Ctx) error {
	var wallet models.Wallet
	userID, _ := c.Locals("userID").(primitive.ObjectID)
	findWallet, err := wallet.FindByUserID(userID)
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: "Wallet does not exist"}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
	fmt.Println("Wallet", findWallet)
	response := &finance.CreateWalletResponse{Success: true, Message: "Wallet Found", Wallet: &findWallet}
	return c.Status(fiber.StatusOK).JSON(response)
}

// arrayFilter := primitive.M{
// 	"userid": userID,
// }
// arrayFilters := options.ArrayFilters{
// 	Filters: []interface{}{arrayFilter},
// }
// count, err := wallet.Count(arrayFilters)
// if err != nil {
// 	response := &handler.ErrorResponse{Success: false, Error: err.Error()}
// 	return c.Status(fiber.StatusCreated).JSON(response)
// }
// fmt.Println("Count: ", count)
// if count > 0 {
// 	response := &handler.ErrorResponse{Success: false, Error: "Wallet already exists"}
// 	return c.Status(fiber.StatusCreated).JSON(response)
// }
