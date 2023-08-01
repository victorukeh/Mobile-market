package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"github.com/victorukeh/mobile-market/pkg/v1/responses"
	"github.com/victorukeh/mobile-market/pkg/v1/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrandController struct{}

func (u *ErrandController) AddErrand(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	var errand models.Errand
	err := c.BodyParser(&errand)
	errand.Initiator = userID
	if err != nil {
		return responses.ErrorResponse(c, err, 400)
	}
	errandService := &services.ErrandService{}
	return errandService.AddErrand(c, errand)
}
