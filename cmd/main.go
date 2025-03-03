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
	"github.com/gofiber/swagger"
)

// @title JWT Auth Microservice API
// @version 1.0
// @description This is a simple authentication microservice using Go Fiber.
// @host localhost:8000
func main() {
	configs := config.NewConfig()

	fmt.Printf("%+v\n", configs)

	utils.InitPassword(configs)
	utils.InitJWT(configs.JWTSecret, configs.JWTExpiry)
	database.Connect(configs)

	app := fiber.New()
	app.Use(middleware.LoggerMiddleware())
	app.Get("/docs/*", swagger.HandlerDefault)

	routers.SetupRouters(app)
	app.Listen(":8000")
}
