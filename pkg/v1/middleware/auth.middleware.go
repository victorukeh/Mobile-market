package middleware

import (
	"fmt"

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

	fmt.Println(claims)

	// sub, ok := claims["sub"].(string)
	// fmt.Println("Claims: ", claims["role"].(String))
	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("userID", claims.Id.Hex())
	c.Set("role", string(claims.Role))
	c.Locals("userID", claims.Id)

	return c.Next()
	// return err
}
