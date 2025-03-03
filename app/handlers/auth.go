package handlers

import (
	"jwt_auth/app/database"
	"jwt_auth/app/models"
	"jwt_auth/app/utils"

	"github.com/gofiber/fiber/v2"
)

// RegisterUser registers a new user
// @Summary Register a new user
// @Description Creates a new user and sets up initial properties
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "User Registration Request"
// @Success 201 {object} any
// @Failure 400 {object} any
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad request"})
	}

	hashedPassword, _ := utils.PWD.HashPassword(req.Password)

	user := models.User{ // TODO: make User.FromRequest function
		Username: req.Username,
		Name:     req.Name,
		Second:   req.Second,
		Password: hashedPassword,
		Properties: models.Properties{ // TODO: remove from register method
			MaxSessions:    5, 
			SessionExpires: 7,
		},
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "User creation failed"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}
