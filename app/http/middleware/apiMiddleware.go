package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrmertkose/go-api/app/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "unauthenticated",
		})

	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "unauthenticated",
		})
	}

	ctx.Locals("id", claims["Id"])

	return ctx.Next()
}
