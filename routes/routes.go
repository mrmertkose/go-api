package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrmertkose/go-api/app/controllers/auth"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)
	app.Post("/logout", auth.Logout)

}
