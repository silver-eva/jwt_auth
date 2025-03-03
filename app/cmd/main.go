package main

import (
	"jwt_auth/app/config"
	"jwt_auth/app/database"
	"jwt_auth/app/middleware"
	"jwt_auth/app/routers"
	"jwt_auth/app/utils"

	_ "jwt_auth/app/docs"

	"github.com/gofiber/fiber/v2"
)

// @title JWT Auth Microservice API
// @version 1.0
// @description This is a simple authentication microservice using Go Fiber.
// @host localhost:8000
func main() {
	configs := config.NewConfig()
	utils.InitPassword(configs)
	database.Connect(configs)

	app := fiber.New()
	app.Use(middleware.LoggerMiddleware())

	routers.SetupRouters(app)
	app.Listen(":8000")
}
