package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/finance"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"github.com/victorukeh/mobile-market/pkg/v1/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinanceController struct{}

// Cash

func (u *FinanceController) CreateCashType(c *fiber.Ctx) error {
	var cashType models.CashType
	err := c.BodyParser(&cashType)
	if err != nil {
		return responses.ErrorResponse(c, err)
	}
	validationErr := validate.Struct(cashType)
	if validationErr != nil {
		return responses.ErrorResponse(c, validationErr)
	}
	result, err := cashType.CreateCashType(cashType)
	if err != nil {
		return responses.ErrorResponse(c, err)
	}
	return responses.WalletCreateResponse(c, result)
}

// Wallets
func (u *FinanceController) GetWallet(c *fiber.Ctx) error {
	var wallet models.Wallet
	userID, _ := c.Locals("userID").(primitive.ObjectID)
	findWallet, err := wallet.FindByUserID(userID)
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: "Wallet does not exist"}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
	fmt.Println("Wallet", findWallet)
	return responses.WalletSuccessResponse(c, findWallet)
}

func (u *FinanceController) SetStatus(c *fiber.Ctx) error {
	var cash models.CashType
	var cashForm models.CashForm
	err := c.BodyParser(&cashForm)
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: err.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	validationErr := validate.Struct(cash)
	if validationErr != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, validationErr.Error())
		response := &handler.ErrorResponse{Success: false, Error: validationErr.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	response := &finance.SetStatusDto{Success: true, Message: "Cash successfully validated"}
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
