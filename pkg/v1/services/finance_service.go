package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"github.com/victorukeh/mobile-market/pkg/v1/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinanceService struct{}

var validate = validator.New()

// Cash
func (u *FinanceService) CreateCashType(c *fiber.Ctx, cashType models.CashType) error {

	err := cashType.FindOneCashType(cashType)
	if err == nil {
		err := errors.New("CashType already exists")
		return responses.ErrorResponse(c, err, 401)
	}
	result, err := cashType.CreateCashType(cashType)
	if err != nil {
		return responses.ErrorResponse(c, err, 500)
	}
	return responses.CashTypeCreateResponse(c, result)
}

func (u *FinanceService) CreateCashGroup(c *fiber.Ctx, cashData models.CashData) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	var cashGroup models.CashGroup
	var cashType models.CashType
	for _, data := range cashData.CashGroup {
		value, err := json.Marshal(&data)
		if err != nil {
			return responses.ErrorResponse(c, err, 400)
		}

		err = json.Unmarshal(value, &cashGroup)
		if err != nil {
			return responses.ErrorResponse(c, err, 400)
		}
		err = cashType.FindCashTypeById(cashGroup.CashTypeID)
		if err != nil {
			fmt.Println(err)
			err := errors.New("CashType not found")
			return responses.ErrorResponse(c, err, 404)
		}
		// Check if Cash Group exists. If it exists then increment
		result, err := cashGroup.FindCashGroup(cashGroup)
		if err != nil {
			cashGroup.UserID = userID
			cashGroup.ID = primitive.NewObjectID()
			validationErr := validate.Struct(cashGroup)
			if validationErr != nil {
				return responses.ErrorResponse(c, validationErr, 400)
			}
			_, err = cashGroup.CreateCashGroup(cashGroup)
		} else {
			var getCashGroup models.CashGroup
			result.Decode(&getCashGroup)
			err = cashGroup.UpdateCashGroup(getCashGroup.ID, getCashGroup.Number+cashGroup.Number)
		}
		if err != nil {
			return responses.ErrorResponse(c, err, 400)
		}
	}
	return responses.CashGroupCreateResponse(c, cashData.CashGroup)
}

// Wallets
func (u *FinanceService) GetWallet(c *fiber.Ctx) error {
	variable := c.Get("email")
	fmt.Println("Variable: ", variable)
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

// func (u *FinanceService) SetStatus(c *fiber.Ctx) error{

// }
