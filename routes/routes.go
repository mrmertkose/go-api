package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrmertkose/go-api/app/controllers/auth"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	app.Post("/api/register", auth.Register)
	app.Post("/api/login", auth.Login)
	app.Post("/api/logout", auth.Logout)

}
