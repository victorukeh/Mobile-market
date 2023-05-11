package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/users"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct{}

// var validate = validator.New()

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response := handler.ErrorResponse{Success: false, Error: "Invalid Mongo ID"}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	result, err := user.FindById(objID)
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: "User does not exist"}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}
	response := &users.GetUser{Success: true, Message: "User Fetched Successfully", User: result}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	var user models.User
	getPage := c.Query("page")
	getLimit := c.Query("limit")
	if getPage == "" {
		getPage = "1"
	}
	if getLimit == "" {
		getLimit = "20"
	}
	fmt.Println(getLimit, getPage)
	page, err := strconv.Atoi(getPage)
	if page < 1{
		page = 1
	}
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(getLimit)
	if err != nil {
		limit = 20
	}
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: err.Error()}
		return c.Status(fiber.StatusOK).JSON(response)
	}

	result, err := user.FindAll(int64(page - 1), int64(limit))
	if err != nil {
		response := &handler.ErrorResponse{Success: false, Error: err.Error()}
		return c.Status(fiber.StatusOK).JSON(response)
	}
	response := &users.GetUsers{Success: true, Message: "Users Fetched Successfully", Limit: limit, Page: page, User: result}
	return c.Status(fiber.StatusOK).JSON(response)
}
