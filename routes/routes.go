package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrmertkose/go-api/app/http/controllers/auth"
	"github.com/mrmertkose/go-api/app/http/middleware"
)

func Setup(app *fiber.App) {

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)

	app.Post("/logout", middleware.Auth, auth.Logout)

}
