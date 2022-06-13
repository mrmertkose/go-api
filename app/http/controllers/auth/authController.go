package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mrmertkose/go-api/app/http/request"
	"github.com/mrmertkose/go-api/app/models"
	"github.com/mrmertkose/go-api/app/utils"
	"github.com/mrmertkose/go-api/config/database"
)

var user models.User

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	// validate request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// check available user
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Wrong credentials",
		})
	}

	//check validation pass
	isValidPass := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValidPass {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Wrong credentials",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   "secret",
	})
}

func Register(ctx *fiber.Ctx) error {
	registerRequest := new(request.RegisterRequest)
	if err := ctx.BodyParser(registerRequest); err != nil {
		return err
	}

	// validate request
	validate := validator.New()
	errValidate := validate.Struct(registerRequest)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := models.User{
		Name:  registerRequest.Name,
		Email: registerRequest.Email,
	}

	// password hash
	password, err := utils.HashingPassword(registerRequest.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}
	newUser.Password = password

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to store data",
			"error":   errValidate.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    newUser,
	})

}

func Logout(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(true)
}
