package auth

import "github.com/gofiber/fiber/v2"

func Register(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(true)
}

func Login(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(true)
}

func Logout(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(true)
}
