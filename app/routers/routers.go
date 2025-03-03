package routers

import (
	"jwt_auth/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	app.Post("/register", handlers.Register)
}