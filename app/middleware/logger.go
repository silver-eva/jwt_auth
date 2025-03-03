package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

// LoggerMiddleware provides detailed request logging
func LoggerMiddleware() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${ip}] ${status} - ${method} ${path} (${latency})\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Local",
	})
}
