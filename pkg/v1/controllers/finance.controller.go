package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/finance"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"github.com/victorukeh/mobile-market/pkg/v1/responses"
	"github.com/victorukeh/mobile-market/pkg/v1/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinanceController struct{}

// Cash
type CashData struct {
	CashGroup []models.CashGroup `json:"cashGroup"`
}

func (u *FinanceController) CreateCashType(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	var cashType models.CashType
	err := c.BodyParser(&cashType)
	cashType.ID = primitive.NewObjectID()
	cashType.UserID = userID
	if err != nil {
		return responses.ErrorResponse(c, err, 400)
	}
	validationErr := validate.Struct(cashType)
	if validationErr != nil {
		return responses.ErrorResponse(c, validationErr, 400)
	}
	financeService := &services.FinanceService{}
	return financeService.CreateCashType(c, cashType)
}

func (u *FinanceController) CreateCashGroup(c *fiber.Ctx) error {
	var cashData models.CashData
	err := c.BodyParser(&cashData)
	if err != nil {
		return responses.ErrorResponse(c, err, 400)
	}

	if err != nil {
		return responses.ErrorResponse(c, err, 400)
	}
	financeService := &services.FinanceService{}
	return financeService.CreateCashGroup(c, cashData)
}

// func (u *FinanceController) GetUserCashGroups(c *fiber.Ctx) error {

// }

// Wallets
func (u *FinanceController) GetWallet(c *fiber.Ctx) error {
	financeService := &services.FinanceService{}
	return financeService.GetWallet(c)
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
