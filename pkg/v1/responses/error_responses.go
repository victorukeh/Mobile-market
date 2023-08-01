package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/handler"
)

func ErrorResponse(c *fiber.Ctx, err error, status int64) error {
	response := &handler.ErrorResponse{Success: false, Error: err.Error()}
	return c.Status(int(status)).JSON(response)
}
