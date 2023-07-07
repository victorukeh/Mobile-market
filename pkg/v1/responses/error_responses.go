package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
)

func ErrorResponse(c *fiber.Ctx, err error) error {
	response := &handler.ErrorResponse{Success: false, Error: err.Error()}
	return c.Status(fiber.ErrBadRequest.Code).JSON(response)
}
