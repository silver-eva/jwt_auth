package main

import (
	"fmt"
	"jwt_auth/config"

	"jwt_auth/database"
	"jwt_auth/middleware"
	"jwt_auth/routers"
	"jwt_auth/utils"

	_ "jwt_auth/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)


// @title JWT Auth API
// @version 2.0
// @description JWT Auth API

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email uOY9m@example.com

// @BasePath /
func main() {
	configs := config.NewConfig()

	// fmt.Printf("%+v\n", configs)

	utils.InitPassword(configs)
	utils.InitJWT(configs.JWTSecret, configs.JWTExpiry)
	database.Connect(configs)

	app := fiber.New()
	app.Use(middleware.LoggerMiddleware())

	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{"status": "ok"})
	})
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	routers.SetupRouters(app)
	app.Listen(fmt.Sprintf(":%d", configs.AppPort))
}
