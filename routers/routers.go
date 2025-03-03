package routers

import (
	"jwt_auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	app.Post("/register", handlers.Register)
}
