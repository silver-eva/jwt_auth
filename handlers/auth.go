package handlers

import (
	"fmt"
	"jwt_auth/database"
	"jwt_auth/models"
	"jwt_auth/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//	@Summary		Register a new user
//	@Description	Creates a new user and sets up initial properties
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RegisterRequest	true	"User Registration Request"
//	@Success		201		{object}	models.ErrorResponse
//	@Failure		400		{object}	models.ErrorResponse
//  @Router			/register [post]
func Register(c *fiber.Ctx) error {

	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Message: "Invalid request body"})
	}

	hashedPassword, _ := utils.PWD.HashPassword(req.Password)

	user := models.User{
		Uname: req.Uname,
		Name:     req.Name,
		Second:   req.Second,
		Password: hashedPassword, 
	}

	tx := database.DB.Begin()
	
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return c.Status(400).JSON(models.ErrorResponse{Message: "User creation failed"})
	}
	if err := tx.Create(&models.Properties{UserID: user.ID}).Error; err != nil {
		tx.Rollback()
		return c.Status(400).JSON(models.ErrorResponse{Message: "User.Properties creation failed"})
	}

	tx.Commit()

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}

// Login authenticates a user based on the provided credentials.
// @Summary User login
// @Description Authenticates a user using username and password, generates access and refresh tokens, and manages user sessions.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.LoginUserRequest true "User login request"
// @Success 200 {object} models.LoggedInUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var req models.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Message: "Invalid request body"})
	}

	// Find user by username
	var user models.User
	if err := database.DB.Where("uname = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Message: "User not found"})
	}

	// Check password
	if !utils.PWD.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(401).JSON(models.ErrorResponse{Message: "Bad credentials"})
	}

	// Get user properties
	var properties models.Properties
	database.DB.Where("user_id = ?", user.ID).First(&properties)

	// Count active sessions
	var activeSessions int64
	database.DB.Model(&models.Session{}).Where("user_id = ? AND active = ?", user.ID, true).Count(&activeSessions)

	// Drop oldest session if max_sessions exceeded
	if activeSessions >= int64(properties.MaxSessions) {
		var oldestSession models.Session
		database.DB.Where("user_id = ? AND active = ?", user.ID, true).
			Order("created ASC").
			First(&oldestSession)

		database.DB.Model(&oldestSession).Update("active", false)
	}

	// Generate tokens
	refreshToken, _ := utils.JWT.GenerateRefreshToken(user.ID, properties.SessionExpires)
	accessToken, _ := utils.JWT.GenerateAccessToken(user.ID)

	// Create new session
	session := models.Session{
		ID:        uuid.New(),
		UserID:    user.ID,
		Created:   time.Now(),
		Refreshed: time.Now(),
		Token:   refreshToken,
		Active:    true,
	}
	if database.DB.Create(&session).Error != nil {
		return c.Status(400).JSON(models.ErrorResponse{Message: "Session creation failed"})
	}

	return c.Status(200).JSON(models.LoggedInUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}


func terminateSession(session models.Session) {
	session.Active = false
	session.Terminated = time.Now()

	database.DB.Save(&session)
}

//	@Summary		Refreshes access token using a valid refresh token
//	@Description	Validates the provided refresh token and generates a new access token.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RefreshTokenRequest	true	"Refresh token request"
//	@Success		200		{object}	models.LoggedInUserResponse
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		401		{object}	models.ErrorResponse
//  @Router			/refresh [post]
func Refresh(c *fiber.Ctx) error {
	var req models.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Message: "Invalid request body"})
	}

	session := models.Session{}
	if err := database.DB.Where("token = ? AND active = ?", req.RefreshToken, true).First(&session).Error; err != nil {
		return c.Status(401).JSON(models.ErrorResponse{Message: "Invalid refresh token"})
	}

	claims, err := utils.JWT.ValidateToken(req.RefreshToken)
	if err != nil {
		terminateSession(session)
		return c.Status(401).JSON(models.ErrorResponse{Message: fmt.Sprintf("Invalid refresh token: %v", err)})
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		terminateSession(session)
		return c.Status(401).JSON(models.ErrorResponse{Message: "Invalid refresh token"})
	}

	accessToken, _ := utils.JWT.GenerateAccessToken(uuid.MustParse(userId))
	refreshToken, _ := utils.JWT.GenerateRefreshToken(uuid.MustParse(userId), 30)

	session.Refreshed = time.Now()
	session.Token = refreshToken
	database.DB.Save(&session)

	return c.Status(200).JSON(models.LoggedInUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}