package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/auth"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
	helper "github.com/victorukeh/mobile-market/pkg/v1/helpers"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct{}

var validate = validator.New()

func (uc *AuthController) Register(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: err.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, validationErr.Error())
		response := &handler.ErrorResponse{Success: false, Error: validationErr.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	password := user.HashPassword(*user.Password)
	user.Password = &password
	count, _ := user.CountUsers("email", *user.Email)
	if count > 0 {
		response := &handler.ErrorResponse{Success: false, Error: "Email has been taken"}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	count, _ = user.CountUsers("phone", *user.Phone)
	if count > 0 {
		response := &handler.ErrorResponse{Success: false, Error: "Phone number is already in use"}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	token, _, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.Role, user.ID)
	user.Confirmation_token = &token

	result, err := user.Create(user)
	if err != nil {
		response := &auth.Response{Success: false, Message: "User item was not created"}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
	response := &auth.RegisterResponse{Success: true, Message: "User Created Successfully", User: result}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// @desc	Login to Mobile Market
// @route	/api/v1/auth/login
// @access	Public
func (uc *AuthController) Login(c *fiber.Ctx) error {
	var user models.User
	var foundUser models.User
	var try auth.LoginForm

	err := c.BodyParser(&try)
	if err != nil {
		response := &auth.Response{Success: false, Message: err.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}

	// Validation
	validationErr := validate.Struct(try)
	if validationErr != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, validationErr.Error())
		response := &auth.Response{Success: false, Message: validationErr.Error()}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}

	result, err := user.FindByEmail(*try.Email, foundUser)

	if err != nil {
		response := &auth.Response{Success: false, Message: "User not found"}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}
	passwordIsValid := user.VerifyPassword(*try.Password, *result.Password)
	if !passwordIsValid {
		response := &auth.Response{Success: false, Message: "Invalid password"}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}

	_, token := helper.SetSignedCookieOrToken(foundUser)
	if err != nil {
		response := &auth.Response{Success: false, Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := &auth.LoginResponse{Success: false, Token: token, Message: "Login Successful", User: result}
	return c.Status(fiber.StatusOK).JSON(response)
}
