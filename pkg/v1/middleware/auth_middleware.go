package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorukeh/mobile-market/pkg/v1/dto/auth"
	helper "github.com/victorukeh/mobile-market/pkg/v1/helpers"
)

func Authentication(c *fiber.Ctx) error {
	clientToken := c.Get("token")
	if clientToken == "" {
		response := &auth.Response{Success: false, Message: "No Authorization header provided"}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}
	claims, err := helper.ValidateToken(clientToken)
	if err != "" {
		response := &auth.Response{Success: false, Message: "Bad Token"}
		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("id", claims.Id.Hex())
	c.Set("role", string(claims.Role))

	return c.Next()
	// return err
}
