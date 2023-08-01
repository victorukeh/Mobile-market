package services

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/models"
	"github.com/victorukeh/mobile-market/pkg/v1/responses"
)

type ErrandService struct{}

func (u *ErrandService) AddErrand(c *fiber.Ctx, errand models.Errand) error {
	_, err := errand.CreateErrand(errand)
	if err == nil {
		err := errors.New("something went wrong")
		return responses.ErrorResponse(c, err, 500)
	}
	return responses.ErrandCreationResponse(c)
}
